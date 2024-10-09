package entity

import "github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"

type Game8Monster struct {
	id              vo.ID
	originMonsterNo vo.No
	url             vo.URL
	scores          []*Game8MonsterScore
}

func NewGame8Monster(
	id vo.ID,
	originMonsterNo vo.No,
	url vo.URL,
	scores []*Game8MonsterScore,
) *Game8Monster {
	return &Game8Monster{
		id:              id,
		originMonsterNo: originMonsterNo,
		url:             url,
		scores:          scores,
	}
}

func (m *Game8Monster) OriginMonsterNo() vo.No {
	return m.originMonsterNo
}

func (m *Game8Monster) URL() vo.URL {
	return m.url
}

func (m *Game8Monster) Scores() []*Game8MonsterScore {
	return m.scores
}

func (m *Game8Monster) SetURL(url vo.URL) {
	m.url = url
}

func (m *Game8Monster) SetScores(scores []*Game8MonsterScore) {
	m.scores = scores
}
