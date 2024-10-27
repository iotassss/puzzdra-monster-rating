package service

import (
	"context"

	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/entity"
	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"
)

type ConvertGame8SourceDataService interface {
	Execute(ctx context.Context, game8MonsterSourceData *entity.Game8MonsterSourceData) (*entity.Game8Monster, error)
}

type ConvertGame8SourceDataSV struct {
	findOriginMonsterByNoService FindOriginMonsterByNoService
}

func NewConvertGame8SourceDataSV(
	findOriginMonsterByNoService FindOriginMonsterByNoService,
) *ConvertGame8SourceDataSV {
	return &ConvertGame8SourceDataSV{
		findOriginMonsterByNoService: findOriginMonsterByNoService,
	}
}

func (s *ConvertGame8SourceDataSV) Execute(
	ctx context.Context,
	game8MonsterSourceData *entity.Game8MonsterSourceData,
) (*entity.Game8Monster, error) {
	originMonster, err := s.findOriginMonsterByNoService.Execute(ctx, game8MonsterSourceData.No())
	if err != nil {
		return nil, err
	}

	scores := make([]*entity.Game8MonsterScore, 0, len(game8MonsterSourceData.Scores()))
	for _, sourceDataScore := range game8MonsterSourceData.Scores() {
		scores = append(scores, entity.NewGame8MonsterScore(
			vo.NewTemporaryID(),
			sourceDataScore.Name(),
			sourceDataScore.LeaderPoint(),
			sourceDataScore.SubLeaderPoint(),
			sourceDataScore.AssistPoint(),
		))
	}

	var originMonsterNo vo.No
	if originMonster == nil {
		originMonsterNo = game8MonsterSourceData.No()
	} else {
		originMonsterNo = originMonster.No()
	}

	fetchedGame8Monster := entity.NewGame8Monster(
		vo.NewTemporaryID(),
		originMonsterNo,
		game8MonsterSourceData.URL(),
		scores,
	)

	return fetchedGame8Monster, nil
}
