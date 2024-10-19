package scraper

import (
	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/entity"
	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"
)

type game8MonsterScoreScrapingResult struct {
	name   string
	leader string
	sub    string
	assist string
}

type game8MonsterScrapingResult struct {
	no     int
	url    string
	scores []*game8MonsterScoreScrapingResult
}

func (r *game8MonsterScrapingResult) toEntity() (*entity.Game8MonsterSourceData, error) {
	no, err := vo.NewNo(r.no)
	if err != nil {
		return nil, err
	}
	url, err := vo.NewURL(r.url)
	if err != nil {
		return nil, err
	}
	scores := make([]*entity.Game8MonsterScoreSourceData, 0, len(r.scores))
	for _, score := range r.scores {
		name, err := vo.NewMonsterName(score.name)
		if err != nil {
			return nil, err
		}
		leader, err := vo.NewGame8MonsterPoint(score.leader)
		if err != nil {
			return nil, err
		}
		sub, err := vo.NewGame8MonsterPoint(score.sub)
		if err != nil {
			return nil, err
		}
		assist, err := vo.NewGame8MonsterPoint(score.assist)
		if err != nil {
			return nil, err
		}
		score := entity.NewGame8MonsterScoreSourceData(
			name,
			leader,
			sub,
			assist,
		)
		scores = append(scores, score)
	}

	return entity.NewGame8MonsterSourceData(
		no,
		url,
		scores,
	), nil
}
