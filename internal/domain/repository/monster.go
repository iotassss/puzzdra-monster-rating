package repository

import (
	"context"

	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/entity"
	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"
)

type MonsterRepository interface {
	FindByNo(ctx context.Context, no vo.No) (*entity.Monster, error)
	Exists(ctx context.Context, no vo.No) (bool, error)
	Save(ctx context.Context, monster *entity.Monster) error
}
