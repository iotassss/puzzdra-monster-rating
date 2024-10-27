package repository

import (
	"context"

	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/entity"
	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"
	"gorm.io/gorm"
)

type MonsterSourceData struct {
	gorm.Model
	No     int    `gorm:"unique;not null"`
	Name   string `gorm:"size:255;not null"`
	BaseNo *int
}

type MonsterSourceDataRepository struct {
	db *gorm.DB
}

func NewMonsterSourceDataRepository(db *gorm.DB) *MonsterSourceDataRepository {
	return &MonsterSourceDataRepository{db: db}
}

func (r *MonsterSourceDataRepository) FindAll(ctx context.Context) ([]*entity.MonsterSourceData, error) {
	var gormMonsterSourceDataList []*MonsterSourceData
	if err := r.db.WithContext(ctx).Find(&gormMonsterSourceDataList).Error; err != nil {
		return nil, err
	}

	var monsterSourceDataList []*entity.MonsterSourceData
	for _, gormMonsterSourceData := range gormMonsterSourceDataList {
		monsterSourceData, err := r.convertToEntity(gormMonsterSourceData)
		if err != nil {
			return nil, err
		}

		monsterSourceDataList = append(monsterSourceDataList, monsterSourceData)
	}

	return monsterSourceDataList, nil
}

func (r *MonsterSourceDataRepository) FindByNo(ctx context.Context, inputNo vo.No) (*entity.MonsterSourceData, error) {
	var gormMonsterSourceData *MonsterSourceData
	if err := r.db.WithContext(ctx).Where("no = ?", inputNo.Value()).First(&gormMonsterSourceData).Error; err != nil {
		return nil, err
	}

	return r.convertToEntity(gormMonsterSourceData)
}

func (r *MonsterSourceDataRepository) Exists(ctx context.Context, no vo.No) (bool, error) {
	var monsterSourceData *MonsterSourceData
	if err := r.db.WithContext(ctx).Where("no = ?", no.Value()).First(&monsterSourceData).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *MonsterSourceDataRepository) Save(ctx context.Context, monsterSourceData *entity.MonsterSourceData) error {
	// check if the record exists
	exists, err := r.Exists(ctx, monsterSourceData.No())
	if err != nil {
		return err
	}

	if exists {
		var gormMonsterSourceData *MonsterSourceData
		if err := r.db.WithContext(ctx).
			Where("no = ?", monsterSourceData.No().Value()).
			First(&gormMonsterSourceData).Error; err != nil {
			return err
		}

		gormMonsterSourceData.No = monsterSourceData.No().Value()
		gormMonsterSourceData.Name = monsterSourceData.Name().Value()
		if monsterSourceData.BaseNo() != nil {
			baseNoValue := monsterSourceData.BaseNo().Value()
			gormMonsterSourceData.BaseNo = &baseNoValue
		} else {
			gormMonsterSourceData.BaseNo = nil
		}

		if err := r.db.WithContext(ctx).Save(&gormMonsterSourceData).Error; err != nil {
			return err
		}
	} else {
		gormMonsterSourceData := MonsterSourceData{
			No:   monsterSourceData.No().Value(),
			Name: monsterSourceData.Name().Value(),
		}
		if monsterSourceData.BaseNo() != nil {
			baseNoValue := monsterSourceData.BaseNo().Value()
			gormMonsterSourceData.BaseNo = &baseNoValue
		}

		if err := r.db.WithContext(ctx).Create(&gormMonsterSourceData).Error; err != nil {
			return err
		}
	}

	return nil
}

func (r *MonsterSourceDataRepository) convertToEntity(gormMonsterSourceData *MonsterSourceData) (*entity.MonsterSourceData, error) {
	if gormMonsterSourceData == nil {
		return nil, nil
	}

	id, err := vo.NewID(int(gormMonsterSourceData.ID))
	if err != nil {
		return nil, err
	}
	no, err := vo.NewNo(gormMonsterSourceData.No)
	if err != nil {
		return nil, err
	}
	name, err := vo.NewMonsterName(gormMonsterSourceData.Name)
	if err != nil {
		return nil, err
	}
	var baseNo *vo.No
	if gormMonsterSourceData.BaseNo != nil {
		baseNoVO, err := vo.NewNo(*gormMonsterSourceData.BaseNo)
		if err != nil {
			return nil, err
		}
		baseNo = &baseNoVO
	}

	monsterSourceData := entity.NewMonsterSourceData(id, no, name, baseNo)

	return monsterSourceData, nil
}
