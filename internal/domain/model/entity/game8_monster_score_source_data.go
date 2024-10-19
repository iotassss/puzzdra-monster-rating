package entity

import "github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"

type Game8MonsterScoreSourceData struct {
	name           vo.MonsterName
	leaderPoint    vo.Game8MonsterPoint
	subLeaderPoint vo.Game8MonsterPoint
	assistPoint    vo.Game8MonsterPoint
}

func NewGame8MonsterScoreSourceData(
	name vo.MonsterName,
	leaderPoint vo.Game8MonsterPoint,
	subLeaderPoint vo.Game8MonsterPoint,
	assistPoint vo.Game8MonsterPoint,
) *Game8MonsterScoreSourceData {
	return &Game8MonsterScoreSourceData{
		name:           name,
		leaderPoint:    leaderPoint,
		subLeaderPoint: subLeaderPoint,
		assistPoint:    assistPoint,
	}
}

func (s *Game8MonsterScoreSourceData) Name() vo.MonsterName {
	return s.name
}

func (s *Game8MonsterScoreSourceData) LeaderPoint() vo.Game8MonsterPoint {
	return s.leaderPoint
}

func (s *Game8MonsterScoreSourceData) SubLeaderPoint() vo.Game8MonsterPoint {
	return s.subLeaderPoint
}

func (s *Game8MonsterScoreSourceData) AssistPoint() vo.Game8MonsterPoint {
	return s.assistPoint
}
