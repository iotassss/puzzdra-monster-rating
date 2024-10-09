package entity

import "github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"

type Monster struct {
	id            vo.ID
	no            vo.No
	name          vo.MonsterName
	originMonster *Monster
}

func NewMonster(
	id vo.ID,
	no vo.No,
	name vo.MonsterName,
	originMonster *Monster,
) *Monster {
	return &Monster{
		id:            id,
		no:            no,
		name:          name,
		originMonster: originMonster,
	}
}

func (m *Monster) No() vo.No {
	return m.no
}

func (m *Monster) Name() vo.MonsterName {
	return m.name
}

func (m *Monster) SetName(name vo.MonsterName) {
	m.name = name
}

func (m *Monster) SetOriginMonster(originMonster *Monster) {
	m.originMonster = originMonster
}
