package usecase

import (
	"context"
	"errors"

	"github.com/iotassss/puzzdra-monster-rating/internal/apperrors"
	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"
	"github.com/iotassss/puzzdra-monster-rating/internal/domain/repository"
)

type GetMonsterRatingPresenter interface {
	Present(monsterRating MonsterRating) error
	PresentError(err error) error
}

type GetMonsterRatingUsecase interface {
	Execute(ctx context.Context, no int) error
}

type GetMonsterRatingUsecaseInteractor struct {
	presenter        GetMonsterRatingPresenter
	monsterRepo      repository.MonsterRepository
	game8MonsterRepo repository.Game8MonsterRepository
}

func NewGetMonsterRatingUsecaseInteractor(
	presenter GetMonsterRatingPresenter,
	monsterRepo repository.MonsterRepository,
	game8MonsterRepo repository.Game8MonsterRepository,
) *GetMonsterRatingUsecaseInteractor {
	return &GetMonsterRatingUsecaseInteractor{
		presenter:        presenter,
		monsterRepo:      monsterRepo,
		game8MonsterRepo: game8MonsterRepo,
	}
}

func (uc *GetMonsterRatingUsecaseInteractor) Execute(ctx context.Context, inputNo int) error {
	no, err := vo.NewNo(inputNo)
	if err != nil {
		err = &apperrors.ErrValidation{Message: "no is invalid", Cause: err}
		return uc.presenter.PresentError(err)
	}

	// get monster
	exist, err := uc.monsterRepo.Exists(ctx, no)
	if err != nil {
		return uc.presenter.PresentError(err)
	}
	if !exist {
		err = &apperrors.ErrNotFound{Message: "monster not found", Cause: errors.New("monster not found")}
		return uc.presenter.PresentError(err)
	}
	monster, err := uc.monsterRepo.FindByNo(ctx, no)
	if err != nil {
		return uc.presenter.PresentError(err)
	}

	// get game8 monster
	var game8MonsterOutput Game8Monster
	exist, err = uc.game8MonsterRepo.Exists(ctx, monster.No())
	if err != nil {
		return uc.presenter.PresentError(err)
	}
	if exist {
		game8Monster, err := uc.game8MonsterRepo.FindByNo(ctx, no)
		if err != nil {
			return uc.presenter.PresentError(err)
		}

		game8MonsterScoresOutput := make([]Game8monsterScore, 0, len(game8Monster.Scores()))
		for _, score := range game8Monster.Scores() {
			game8MonsterScoresOutput = append(game8MonsterScoresOutput, Game8monsterScore{
				Name:           score.Name().Value(),
				LeaderPoint:    score.LeaderPoint().Value(),
				SubLeaderPoint: score.SubLeaderPoint().Value(),
				AssistPoint:    score.AssistPoint().Value(),
			})
		}

		url := game8Monster.URL().Value()
		game8MonsterOutput = Game8Monster{
			URL:             url.String(),
			OriginMonsterNo: game8Monster.OriginMonsterNo().Value(),
			Scores:          game8MonsterScoresOutput,
		}
	}

	// presentation data
	monsterRating := MonsterRating{
		No:           monster.No().Value(),
		Name:         monster.Name().Value(),
		Game8Monster: game8MonsterOutput,
	}

	return uc.presenter.Present(monsterRating)
}
