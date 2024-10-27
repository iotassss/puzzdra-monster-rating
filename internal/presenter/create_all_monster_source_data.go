package presenter

import "fmt"

type CreateAllMonsterSourceDataPresenter struct {
}

func NewCreateAllMonsterSourceDataPresenter() *CreateAllMonsterSourceDataPresenter {
	return &CreateAllMonsterSourceDataPresenter{}
}

func (p *CreateAllMonsterSourceDataPresenter) Present() error {
	fmt.Println("CreateAllMonsterSourceData is done.")
	return nil
}
