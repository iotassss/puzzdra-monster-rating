package loader

import (
	"context"
	"encoding/json"
	"os"

	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/entity"
	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"
)

type Evolution struct {
	BaseNo *int `json:"baseNo"`
}

type MonsterSourceData struct {
	ID        int       `json:"id"`
	No        int       `json:"no"`
	Name      string    `json:"name"`
	Evolution Evolution `json:"evolution"`
}

type MonsterSourceDataLoader struct {
	monsterJSONFilePath string
}

func NewMonsterSourceDataLoader(monsterJSONFilePath string) *MonsterSourceDataLoader {
	return &MonsterSourceDataLoader{
		monsterJSONFilePath: monsterJSONFilePath,
	}
}

/*
指定のjsonファイルから全てのモンスター情報を読み込み、entity.MonsterSourceDataのスライスとして返す。

jsonファイルの形式は以下のフィールドを持つオブジェクトの配列であることを期待する。

[

	{
		"no": 1,
		"name": "モンスター1",
		"evolution": {
			"baseNo": 2
		}
	},
	{
		...
	},
	...

]
*/
func (l *MonsterSourceDataLoader) LoadAll(ctx context.Context) ([]*entity.MonsterSourceData, error) {
	// filePathからデータをstreamで1オブジェクトずつ読み込む
	// 1オブジェクトずつentity.MonsterSourceDataに変換して返す

	file, err := os.Open(l.monsterJSONFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	// JSONが配列の形式であるか確認する
	if _, err := decoder.Token(); err != nil {
		return nil, err
	}

	var monsters []*entity.MonsterSourceData

	// JSON配列の各要素を順番にデコードする
	for decoder.More() {
		var jsonMonster MonsterSourceData
		if err := decoder.Decode(&jsonMonster); err != nil {
			return nil, err
		}

		monster, err := l.convertToEntity(jsonMonster)
		if err != nil {
			return nil, err
		}

		monsters = append(monsters, monster)
	}

	// JSON配列の終了を確認する
	if _, err := decoder.Token(); err != nil {
		return nil, err
	}

	return monsters, nil
}

func (l *MonsterSourceDataLoader) convertToEntity(jsonMonster MonsterSourceData) (*entity.MonsterSourceData, error) {
	no, err := vo.NewNo(jsonMonster.No)
	if err != nil {
		return nil, err
	}
	name, err := vo.NewMonsterName(jsonMonster.Name)
	if err != nil {
		return nil, err
	}
	var baseNo *vo.No
	if jsonMonster.Evolution.BaseNo != nil {
		tmpBaseNo, err := vo.NewNo(*jsonMonster.Evolution.BaseNo)
		if err != nil {
			return nil, err
		}
		baseNo = &tmpBaseNo
	}

	return entity.NewMonsterSourceData(
		vo.NewTemporaryID(),
		no,
		name,
		baseNo,
	), nil
}
