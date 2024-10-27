package presenter

import "fmt"

type CreateAllMonstersPresenter struct {
}

func NewCreateAllMonstersPresenter() *CreateAllMonstersPresenter {
	return &CreateAllMonstersPresenter{}
}

func (p *CreateAllMonstersPresenter) Present() error {
	fmt.Println("CreateAllMonsters is done.")
	return nil
}
