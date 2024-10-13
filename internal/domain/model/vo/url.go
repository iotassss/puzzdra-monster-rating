package vo

import (
	"fmt"
	"net/url"
)

type ErrURLValidation struct {
	Err error
}

func (e ErrURLValidation) Error() string {
	return fmt.Sprintf("invalid URL: %s", e.Err.Error())
}

type URL struct {
	url url.URL
}

// NewURL コンストラクタ：文字列からURLを作成し、構文が正しいかをチェック
func NewURL(rawURL string) (URL, error) {
	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return URL{}, ErrURLValidation{Err: err}
	}
	return URL{url: *parsedURL}, nil
}

// 現状はURLの値をそのまま返す
func (u URL) Value() url.URL {
	return u.url
}
