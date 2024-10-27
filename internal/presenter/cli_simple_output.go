package presenter

import "fmt"

type CLISimpleOutputPresenter struct {
	message string
}

func NewCLISimpleOutputPresenter(message string) *CLISimpleOutputPresenter {
	return &CLISimpleOutputPresenter{
		message: message,
	}
}

func (p *CLISimpleOutputPresenter) Present() error {
	fmt.Println(p.message)
	return nil
}
