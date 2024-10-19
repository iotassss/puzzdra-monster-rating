package scraper

import (
	"context"
	"log/slog"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/entity"
	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"
)

const (
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36"
)

type Game8MonsterScraper struct {
	logger        *slog.Logger
	timeoutSecond *int
}

func NewGame8MonsterScraper(
	logger *slog.Logger,
	timeout *int,
) *Game8MonsterScraper {
	if timeout == nil {
		defaultTimeout := 5
		timeout = &defaultTimeout
	}

	return &Game8MonsterScraper{
		logger:        logger,
		timeoutSecond: timeout,
	}
}

func (s *Game8MonsterScraper) Fetch(ctx context.Context, url vo.URL) (*entity.Game8MonsterSourceData, error) {
	result := &game8MonsterScrapingResult{}

	c := colly.NewCollector(
		colly.UserAgent(userAgent),
	)
	c.SetRequestTimeout(time.Duration(*s.timeoutSecond) * time.Second)

	c.OnError(func(r *colly.Response, err error) {
		s.logger.Error("Scraping error", slog.Any("error", err))
	})

	// No取得
	c.OnHTML("h3", func(e *colly.HTMLElement) {
		// h3要素が「のステータス」を含むことを確認
		if !strings.Contains(e.Text, "のステータス") {
			return
		}

		// h3タグの次のsiblingがtableかどうかを確認
		table := e.DOM.Next()
		if goquery.NodeName(table) != "table" {
			return
		}

		// tableの最初の行のth要素のテキストを確認
		thText := table.Find("tr:first-child th").Text()
		if thText == "" {
			return
		}

		// 【No.xxx】モンスター名 の形式からxxx部分の数値を正規表現で抽出
		re := regexp.MustCompile(`【No\.\s*(\d+)】`)
		match := re.FindStringSubmatch(thText)
		if len(match) < 2 {
			return
		}
		noStr := match[1]

		// 保存
		no, err := strconv.Atoi(noStr)
		if err != nil {
			s.logger.Error("Failed to convert string to int", slog.String("noStr", noStr), slog.Any("error", err))
			return
		}

		result.no = no
	})

	// 点数取得
	c.OnHTML("table", func(e *colly.HTMLElement) {
		var isScoreTable bool
		var name string
		var isPattern2 bool
		e.ForEach("tr", func(index int, row *colly.HTMLElement) {
			// 最初の行に「リーダー評価」「サブ評価」「アシスト評価」の文字列があるか確認
			isScoreTableHeader := strings.Contains(row.Text, "リーダー") &&
				strings.Contains(row.Text, "サブ") &&
				strings.Contains(row.Text, "アシスト")

			if index == 0 && isScoreTableHeader {
				if strings.Contains(row.Text, "リーダー評価") {
					isPattern2 = true
				}
				isScoreTable = true
				return
			}

			// リーダー評価などが確認されたら次の行から点数を取得
			if isScoreTable {
				var leader, sub, assist string
				if isPattern2 {
					table := row.DOM.Parent().Parent()
					p := table.Prev()
					h2 := p.Prev()

					if goquery.NodeName(h2) != "h2" || !strings.Contains(h2.Text(), "の評価") {
						return
					}

					name = strings.Replace(h2.Text(), "の評価", "", -1)
					name = strings.Replace(name, "と使い道", "", -1)
					leader = row.ChildText("td:nth-of-type(1)")
					leader = strings.TrimSuffix(leader, "点 / 9.9点")
					sub = row.ChildText("td:nth-of-type(2)")
					sub = strings.TrimSuffix(sub, "点 / 9.9点")
					assist = row.ChildText("td:nth-of-type(3)")
					assist = strings.TrimSuffix(assist, "点 / 9.9点")
				} else {
					name = row.ChildText("td:nth-of-type(1)")
					leader = row.ChildText("td:nth-of-type(2)")
					sub = row.ChildText("td:nth-of-type(3)")
					assist = row.ChildText("td:nth-of-type(4)")
				}

				result.scores = append(result.scores, &game8MonsterScoreScrapingResult{
					name:   name,
					leader: leader,
					sub:    sub,
					assist: assist,
				})
			}
		})
	})

	err := c.Visit(url.Value().String())
	if err != nil {
		s.logger.Error("Failed to visit page", slog.String("url", url.Value().String()), slog.Any("error", err))
		return nil, err
	}

	game8Monster, err := result.toEntity()
	if err != nil {
		s.logger.Error("Failed to convert scraping result to entity", slog.Any("error", err))
		return nil, err
	}

	return game8Monster, nil
}
