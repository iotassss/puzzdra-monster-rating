package scraper

import (
	"context"
	"fmt"
	"os"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type Game8MonsterURLScraper struct {
	timeoutSecond int
	userAgent     string
	outputFile    string
	debug         bool
}

type Game8MonsterURLScraperConfig struct {
	TimeoutSecond int
	UserAgent     string
	OutputFile    string
	Debug         bool
}

func NewGame8MonsterURLScraper(config *Game8MonsterURLScraperConfig) *Game8MonsterURLScraper {
	timeoutSecond := config.TimeoutSecond
	if timeoutSecond <= 0 {
		timeoutSecond = 5
	}

	return &Game8MonsterURLScraper{
		timeoutSecond: timeoutSecond,
		userAgent:     config.UserAgent,
		outputFile:    config.OutputFile,
		debug:         config.Debug,
	}
}

func (scraper *Game8MonsterURLScraper) Fetch(ctx context.Context) error {
	file, err := os.Create(scraper.outputFile)
	if err != nil {
		if scraper.debug {
			fmt.Println("ファイル作成に失敗しました:", err)
		}
		return err
	}
	defer file.Close()

	c := colly.NewCollector()

	c.OnHTML("h3", func(e *colly.HTMLElement) {
		// 副属性「*」のパターンを正規表現に変換
		pattern := `^副属性「.*」のキャラ評価$`
		re := regexp.MustCompile(pattern)
		if !re.MatchString(e.Text) {
			return
		}

		// h3タグの次にあるtableを取得
		nextSibling := e.DOM.Next()
		if goquery.NodeName(nextSibling) != "table" {
			return
		}

		// table内のリンクを取得
		nextSibling.Find("tr").Each(func(i int, selection *goquery.Selection) {
			selection.Find("td a").Each(func(j int, a *goquery.Selection) {
				href, exists := a.Attr("href")
				if exists {
					file.WriteString(href + "\n")
					if scraper.debug {
						fmt.Printf("[%s](%s),", a.Text(), href)
					}
				}
			})
		})
	})

	// モンスター一覧ページのURL
	monsterListUrls := []string{
		"https://game8.jp/pazudora/24173", // 火属性
		"https://game8.jp/pazudora/24241", // 水属性
		"https://game8.jp/pazudora/24242", // 木属性
		"https://game8.jp/pazudora/24243", // 光属性
		"https://game8.jp/pazudora/24236", // 闇属性
	}

	for _, url := range monsterListUrls {
		c.Visit(url)
	}

	return nil
}
