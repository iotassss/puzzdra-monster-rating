package presenter

import "fmt"

type CreateAllGame8MonstersPresenter struct {
}

func NewCreateAllGame8MonstersPresenter() *CreateAllGame8MonstersPresenter {
	return &CreateAllGame8MonstersPresenter{}
}

func (p *CreateAllGame8MonstersPresenter) Present() error {
	fmt.Println("CreateAllGame8Monsters is done.")
	return nil
}
