package repository

import (
	"context"

	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/entity"
	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"
)

type Game8MonsterRepository interface {
	FindByNo(ctx context.Context, no vo.No) (*entity.Game8Monster, error)
	Exists(ctx context.Context, no vo.No) (bool, error)
	Save(ctx context.Context, game8Monster *entity.Game8Monster) error
}

type Game8MonsterURLLoader interface {
	LoadAll(ctx context.Context) ([]vo.URL, error)
}

type Game8MonsterFetcher interface {
	Fetch(ctx context.Context, url vo.URL) (*entity.Game8Monster, error)
}
