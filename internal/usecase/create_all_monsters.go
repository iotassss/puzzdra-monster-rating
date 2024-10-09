package usecase

import (
	"context"

	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/entity"
	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"
	"github.com/iotassss/puzzdra-monster-rating/internal/domain/repository"
	"github.com/iotassss/puzzdra-monster-rating/internal/domain/service"
)

type CreateAllMonstersPresenter interface {
	Present() error
}

type CreateAllMonstersUsecase interface {
	Execute(ctx context.Context) error
}

type CreateAllMonstersUsecaseInteractor struct {
	presenter                    CreateAllMonstersPresenter
	monsterRepo                  repository.MonsterRepository
	monsterSourceDataRepo        repository.MonsterSourceDataRepository
	findOriginMonsterByNoService service.FindOriginMonsterByNoService
}

func NewCreateAllMonstersUsecaseInteractor(
	presenter CreateAllMonstersPresenter,
	monsterRepo repository.MonsterRepository,
	monsterSourceDataRepo repository.MonsterSourceDataRepository,
	findOriginMonsterByNoService service.FindOriginMonsterByNoService,
) *CreateAllMonstersUsecaseInteractor {
	return &CreateAllMonstersUsecaseInteractor{
		presenter:                    presenter,
		monsterRepo:                  monsterRepo,
		monsterSourceDataRepo:        monsterSourceDataRepo,
		findOriginMonsterByNoService: findOriginMonsterByNoService,
	}
}

func (uc *CreateAllMonstersUsecaseInteractor) Execute(ctx context.Context) error {
	// ＊全てのMonsterSourceDataをメモリに展開しているため注意
	monsterSourceDataList, err := uc.monsterSourceDataRepo.FindAll(ctx)
	if err != nil {
		return err
	}

	for _, monsterSourceData := range monsterSourceDataList {
		var originMonster *entity.Monster
		originMonster, err = uc.findOriginMonsterByNoService.Execute(ctx, monsterSourceData.No())
		if err != nil {
			return err
		}

		exists, err := uc.monsterRepo.Exists(ctx, monsterSourceData.No())
		if err != nil {
			return err
		}

		if exists {
			monster, err := uc.monsterRepo.FindByNo(ctx, monsterSourceData.No())
			if err != nil {
				return err
			}
			monster.SetName(monsterSourceData.Name())
			monster.SetOriginMonster(originMonster)
			err = uc.monsterRepo.Save(ctx, monster)
			if err != nil {
				return err
			}
		} else {
			monster := entity.NewMonster(
				vo.NewTemporaryID(),
				monsterSourceData.No(),
				monsterSourceData.Name(),
				originMonster,
			)
			err = uc.monsterRepo.Save(ctx, monster)
			if err != nil {
				return err
			}
		}
	}

	return uc.presenter.Present()
}
