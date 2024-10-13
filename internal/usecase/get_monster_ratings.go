package usecase

import (
	"context"

	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"
	"github.com/iotassss/puzzdra-monster-rating/internal/domain/repository"
)

type GetMonsterRatingsPresenter interface {
	Present(monsterRating MonsterRating) error
}

type GetAllSiteMonsterRatingsUsecase interface {
	Execute(ctx context.Context, no int) error
}

type GetMonsterRatingsUsecaseInteractor struct {
	presenter        GetMonsterRatingsPresenter
	monsterRepo      repository.MonsterRepository
	game8MonsterRepo repository.Game8MonsterRepository
}

func NewGetMonsterRatingsUsecaseInteractor(
	presenter GetMonsterRatingsPresenter,
	monsterRepo repository.MonsterRepository,
	game8MonsterRepo repository.Game8MonsterRepository,
) *GetMonsterRatingsUsecaseInteractor {
	return &GetMonsterRatingsUsecaseInteractor{
		presenter:        presenter,
		monsterRepo:      monsterRepo,
		game8MonsterRepo: game8MonsterRepo,
	}
}

func (uc *GetMonsterRatingsUsecaseInteractor) Execute(ctx context.Context, inputNo int) error {
	no, err := vo.NewNo(inputNo)
	if err != nil {
		return err
	}
	monster, err := uc.monsterRepo.FindByNo(ctx, no)
	if err != nil {
		return err
	}
	game8Monster, err := uc.game8MonsterRepo.FindByNo(ctx, no)
	if err != nil {
		return err
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
	game8MonsterOutput := Game8Monster{
		URL:             url.String(),
		OriginMonsterNo: game8Monster.OriginMonsterNo().Value(),
		Scores:          game8MonsterScoresOutput,
	}

	monsterRating := MonsterRating{
		No:           monster.No().Value(),
		Name:         monster.Name().Value(),
		Game8Monster: game8MonsterOutput,
	}

	return uc.presenter.Present(monsterRating)
}
