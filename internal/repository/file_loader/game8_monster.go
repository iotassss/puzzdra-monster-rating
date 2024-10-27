package loader

import (
	"bufio"
	"context"
	"os"

	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"
)

type Game8MonsterURLLoader struct {
	urlListFilePath string
}

func NewGame8MonsterURLLoader(urlListFilePath string) *Game8MonsterURLLoader {
	return &Game8MonsterURLLoader{
		urlListFilePath: urlListFilePath,
	}
}

/*
指定のファイルからURLの一覧を読み込み、vo.URLのスライスとして返す。

ファイルの形式は以下の形式であることを期待する。

https://example.com/monster1
https://example.com/monster2
https://example.com/monster3
*/
func (l *Game8MonsterURLLoader) LoadAll(ctx context.Context) ([]vo.URL, error) {
	// filePathからデータをstreamで1行ずつ読み込む
	// 1行ずつvo.URLに変換してスライスに追加して返す

	file, err := os.Open(l.urlListFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 1回目の読み込みで行数をカウント
	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	// ファイルの先頭に戻る
	if _, err := file.Seek(0, 0); err != nil {
		return nil, err
	}

	// スライスの初期キャパシティを指定して作成
	urls := make([]vo.URL, 0, lineCount)

	// 2回目の読み込みでURLを処理
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		url, err := vo.NewURL(scanner.Text())
		if err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	return urls, nil
}
