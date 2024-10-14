package usecase

import (
	"context"

	"github.com/iotassss/puzzdra-monster-rating/internal/domain/repository"
)

type CreateAllGame8MonstersPresenter interface {
	Present() error
}

type CreateAllGame8MonstersUsecase interface {
	Execute(ctx context.Context) error
}

type CreateAllGame8MonstersUsecaseInteractor struct {
	presenter             CreateAllGame8MonstersPresenter
	game8MonsterRepo      repository.Game8MonsterRepository
	game8MonsterURLLoader repository.Game8MonsterURLLoader
	game8MonsterFetcher   repository.Game8MonsterFetcher
}

func NewCreateAllGame8MonstersUsecaseInteractor(
	presenter CreateAllGame8MonstersPresenter,
	game8MonsterRepo repository.Game8MonsterRepository,
	game8MonsterURLLoader repository.Game8MonsterURLLoader,
	game8MonsterFetcher repository.Game8MonsterFetcher,
) *CreateAllGame8MonstersUsecaseInteractor {
	return &CreateAllGame8MonstersUsecaseInteractor{
		presenter:             presenter,
		game8MonsterRepo:      game8MonsterRepo,
		game8MonsterURLLoader: game8MonsterURLLoader,
		game8MonsterFetcher:   game8MonsterFetcher,
	}
}

func (uc *CreateAllGame8MonstersUsecaseInteractor) Execute(ctx context.Context) error {
	urls, err := uc.game8MonsterURLLoader.LoadAll(ctx)
	if err != nil {
		return err
	}

	for _, url := range urls {
		fetchedGame8Monster, err := uc.game8MonsterFetcher.Fetch(ctx, url)
		if err != nil {
			return err
		}

		exists, err := uc.game8MonsterRepo.Exists(ctx, fetchedGame8Monster.OriginMonsterNo())
		if err != nil {
			return err
		}

		if exists {
			existingGame8Monster, err := uc.game8MonsterRepo.FindByNo(ctx, fetchedGame8Monster.OriginMonsterNo())
			if err != nil {
				return err
			}
			existingGame8Monster.SetURL(fetchedGame8Monster.URL())
			existingGame8Monster.SetScores(fetchedGame8Monster.Scores())
			err = uc.game8MonsterRepo.Save(ctx, existingGame8Monster)
			if err != nil {
				return err
			}
		} else {
			err = uc.game8MonsterRepo.Save(ctx, fetchedGame8Monster)
			if err != nil {
				return err
			}
		}
	}

	return uc.presenter.Present()
}
