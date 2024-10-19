package service

import (
	"context"

	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/entity"
	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"
)

type ConvertGame8SourceData struct {
	findOriginMonsterByNoService FindOriginMonsterByNo
}

func NewConvertGame8SourceData() *ConvertGame8SourceData {
	return &ConvertGame8SourceData{}
}

func (s *ConvertGame8SourceData) Execute(
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
	fetchedGame8Monster := entity.NewGame8Monster(
		vo.NewTemporaryID(),
		originMonster.No(),
		game8MonsterSourceData.URL(),
		scores,
	)

	return fetchedGame8Monster, nil
}
