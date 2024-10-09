package entity

import "github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"

type Game8MonsterScore struct {
	id             vo.ID
	name           vo.MonsterName
	leaderPoint    vo.Game8MonsterPoint
	subLeaderPoint vo.Game8MonsterPoint
	assistPoint    vo.Game8MonsterPoint
}

func NewGame8MonsterScore(
	id vo.ID,
	name vo.MonsterName,
	leaderPoint vo.Game8MonsterPoint,
	subLeaderPoint vo.Game8MonsterPoint,
	assistPoint vo.Game8MonsterPoint,
) *Game8MonsterScore {
	return &Game8MonsterScore{
		id:             id,
		name:           name,
		leaderPoint:    leaderPoint,
		subLeaderPoint: subLeaderPoint,
		assistPoint:    assistPoint,
	}
}

func (s *Game8MonsterScore) Name() vo.MonsterName {
	return s.name
}

func (s *Game8MonsterScore) LeaderPoint() vo.Game8MonsterPoint {
	return s.leaderPoint
}

func (s *Game8MonsterScore) SubLeaderPoint() vo.Game8MonsterPoint {
	return s.subLeaderPoint
}

func (s *Game8MonsterScore) AssistPoint() vo.Game8MonsterPoint {
	return s.assistPoint
}
