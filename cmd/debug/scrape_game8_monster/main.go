package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func loadDataIntoMemory() ([]string, error) {
	// 入力ファイルを開く
	inputFile, err := os.Open("data/game8_monster_urls.txt")
	if err != nil {
		return nil, fmt.Errorf("入力ファイルを開けませんでした: %v", err)
	}
	defer inputFile.Close()

	// 全件URLをメモリに展開
	var urls []string
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ファイル読み込み中にエラーが発生しました: %v", err)
	}

	return urls, nil
}

// DOMの名前を取得するヘルパー関数
func goqueryNodeName(node *goquery.Selection) string {
	return goquery.NodeName(node)
}

type Score struct {
	Name   string
	Leader string
	Sub    string
	Assist string
}

type MonsterRating struct {
	URL    string
	No     string
	Scores []Score
}

type App struct {
	Collector     *colly.Collector
	MonsterRating *MonsterRating
}

func (app *App) setupCollector() {
	// エラーハンドリング
	app.Collector.OnError(func(r *colly.Response, err error) {
		log.Printf("スクレイピングエラー: %v", err)
	})

	// No取得
	app.Collector.OnHTML("h3", func(e *colly.HTMLElement) {
		// h3要素が「のステータス」を含むことを確認
		if !strings.Contains(e.Text, "のステータス") {
			return
		}

		// h3タグの次のsiblingがtableかどうかを確認
		table := e.DOM.Next()
		if goqueryNodeName(table) != "table" {
			return
		}

		// tableの最初の行のth要素のテキストを確認
		thText := table.Find("tr:first-child th").Text()
		if thText == "" {
			return
		}

		// 【No.xxx】モンスター名 の形式からxxx部分の数値を正規表現で抽出
		re := regexp.MustCompile(`【No\.(\d+)】`)
		match := re.FindStringSubmatch(thText)
		if len(match) < 2 {
			return
		}

		// xxxの部分を取得して保存
		if app.MonsterRating == nil {
			return
		}
		app.MonsterRating.No = match[1]
	})

	// 点数取得
	app.Collector.OnHTML("table", func(e *colly.HTMLElement) {
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

					if goqueryNodeName(h2) != "h2" || !strings.Contains(h2.Text(), "の評価") {
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

				app.MonsterRating.Scores = append(app.MonsterRating.Scores, Score{
					Name:   name,
					Leader: leader,
					Sub:    sub,
					Assist: assist,
				})
			}
		})
	})
}

func (app *App) scrapePage(url string, monsterRating *MonsterRating) error {
	app.MonsterRating = monsterRating
	err := app.Collector.Visit(url)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// メモリに展開したURLを取得
	urls, err := loadDataIntoMemory()
	if err != nil {
		log.Fatalf("データ読み込み中にエラーが発生しました: %v", err)
	}

	// 出力ファイルを作成または開く
	outputFile, err := os.Create("data/debug/game8_monsters_test_access.txt")
	if err != nil {
		log.Fatalf("出力ファイルを作成できませんでした: %v", err)
	}
	defer outputFile.Close()
	writer := bufio.NewWriter(outputFile)

	// 新しいCollyコレクターを作成
	app := App{Collector: colly.NewCollector()}
	app.setupCollector()

	// メモリに展開したURLを処理
	for i, url := range urls {
		fmt.Printf("アクセス中: %s %s\n", fmt.Sprintf("%d", i), url)

		var monsterRating MonsterRating
		err := app.scrapePage(url, &monsterRating)
		if err != nil {
			log.Printf("スクレイピングエラー: %v", err)
			continue
		}

		// 結果を出力ファイルに書き込む
		// 処理番号、url, no, scoresをフォーマットして書き込む
		_, err = writer.WriteString(fmt.Sprintf("%d\n%s\n%s\n%v\n", i, url, monsterRating.No, monsterRating.Scores))
		// 標準出力にも出力
		fmt.Printf("%d\t%s\t%s\t%v\n", i, url, monsterRating.No, monsterRating.Scores)

		// app.MonsterRatingを初期化
		app.MonsterRating = nil

		if (i+1)%10 == 0 {
			writer.Flush()
		}

		// デバッグ用に20件だけ処理
		if (i + 1) == 20 {
			break
		}

		// サイトに負荷をかけないため一定時間待機
		time.Sleep(2 * time.Second)
	}

	// 出力ファイルのバッファをフラッシュ
	writer.Flush()
}
