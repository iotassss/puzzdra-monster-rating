package entity

import "github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"

type Game8MonsterSourceData struct {
	no     vo.No
	url    vo.URL
	scores []*Game8MonsterScoreSourceData
}

func NewGame8MonsterSourceData(
	no vo.No,
	url vo.URL,
	scores []*Game8MonsterScoreSourceData,
) *Game8MonsterSourceData {
	return &Game8MonsterSourceData{
		no:     no,
		url:    url,
		scores: scores,
	}
}

func (m *Game8MonsterSourceData) No() vo.No {
	return m.no
}

func (m *Game8MonsterSourceData) URL() vo.URL {
	return m.url
}

func (m *Game8MonsterSourceData) Scores() []*Game8MonsterScoreSourceData {
	return m.scores
}

func (m *Game8MonsterSourceData) SetURL(url vo.URL) {
	m.url = url
}

func (m *Game8MonsterSourceData) SetScores(scores []*Game8MonsterScoreSourceData) {
	m.scores = scores
}
