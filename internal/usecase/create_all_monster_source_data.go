package usecase

import (
	"context"

	"github.com/iotassss/puzzdra-monster-rating/internal/domain/repository"
)

type CreateAllMonsterSourceDataPresenter interface {
	Present() error
}

type CreateAllMonsterSourceDataUsecase interface {
	Execute(ctx context.Context) error
}

type CreateAllMonsterSourceDataUsecaseInteractor struct {
	presenter               CreateAllMonsterSourceDataPresenter
	monsterSourceDataRepo   repository.MonsterSourceDataRepository
	monsterSourceDataLoader repository.MonsterSourceDataLoader
}

func NewCreateAllMonsterSourceDataUsecaseInteractor(
	presenter CreateAllMonsterSourceDataPresenter,
	monsterSourceDataRepo repository.MonsterSourceDataRepository,
	monsterSourceDataLoader repository.MonsterSourceDataLoader,
) *CreateAllMonsterSourceDataUsecaseInteractor {
	return &CreateAllMonsterSourceDataUsecaseInteractor{
		presenter:               presenter,
		monsterSourceDataRepo:   monsterSourceDataRepo,
		monsterSourceDataLoader: monsterSourceDataLoader,
	}
}

func (uc *CreateAllMonsterSourceDataUsecaseInteractor) Execute(ctx context.Context) error {
	// *caution: all MonsterSourceData is loaded into memory
	rawMonsterSourceDataList, err := uc.monsterSourceDataLoader.LoadAll(ctx)
	if err != nil {
		return err
	}

	for _, rawMonsterSourceData := range rawMonsterSourceDataList {
		err = uc.monsterSourceDataRepo.Save(ctx, rawMonsterSourceData)
		if err != nil {
			return err
		}
	}

	return uc.presenter.Present()
}
