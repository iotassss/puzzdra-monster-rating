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
		exists, err := uc.monsterSourceDataRepo.Exists(ctx, rawMonsterSourceData.No())
		if err != nil {
			return err
		}

		if exists {
			monsterSourceData, err := uc.monsterSourceDataRepo.FindByNo(ctx, rawMonsterSourceData.No())
			if err != nil {
				return err
			}
			monsterSourceData.SetName(rawMonsterSourceData.Name())
			monsterSourceData.SetBaseNo(rawMonsterSourceData.BaseNo())
			err = uc.monsterSourceDataRepo.Save(ctx, monsterSourceData)
			if err != nil {
				return err
			}
		} else {
			err = uc.monsterSourceDataRepo.Save(ctx, rawMonsterSourceData)
			if err != nil {
				return err
			}
		}
	}

	return uc.presenter.Present()
}
