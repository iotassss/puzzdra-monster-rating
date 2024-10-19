package scraper

import "fmt"

type ErrGame8MonsterScraping struct {
	Err error
	Msg string
}

func (e *ErrGame8MonsterScraping) Error() string {
	return fmt.Sprintf("%s: %v", e.Msg, e.Err)
}
