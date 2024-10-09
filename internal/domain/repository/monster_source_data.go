package repository

import (
	"context"

	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/entity"
	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"
)

type MonsterSourceDataRepository interface {
	FindAll(ctx context.Context) ([]*entity.MonsterSourceData, error)
	FindByNo(ctx context.Context, no vo.No) (*entity.MonsterSourceData, error)
	Exists(ctx context.Context, no vo.No) (bool, error)
	Save(ctx context.Context, monsterSourceData *entity.MonsterSourceData) error
}

type MonsterSourceDataLoader interface {
	LoadAll(ctx context.Context) ([]*entity.MonsterSourceData, error)
}
