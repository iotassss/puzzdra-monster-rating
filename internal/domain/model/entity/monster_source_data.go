package entity

import "github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"

/*
monster.jsonから読み込んだデータを保持する構造体
このデータをもとにMonsterを生成する
*/
type MonsterSourceData struct {
	id     vo.ID
	no     vo.No
	name   vo.MonsterName
	baseNo *vo.No
}

func NewMonsterSourceData(
	id vo.ID,
	no vo.No,
	name vo.MonsterName,
	baseNo *vo.No,
) *MonsterSourceData {
	return &MonsterSourceData{
		id:     id,
		no:     no,
		name:   name,
		baseNo: baseNo,
	}
}

func (m *MonsterSourceData) No() vo.No {
	return m.no
}

func (m *MonsterSourceData) Name() vo.MonsterName {
	return m.name
}

func (m *MonsterSourceData) BaseNo() *vo.No {
	return m.baseNo
}

func (m *MonsterSourceData) SetName(name vo.MonsterName) {
	m.name = name
}

func (m *MonsterSourceData) SetBaseNo(baseNo *vo.No) {
	m.baseNo = baseNo
}
