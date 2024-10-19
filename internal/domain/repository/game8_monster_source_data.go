package repository

import (
	"context"

	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/entity"
	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"
)

type Game8MonsterSourceDataFetcher interface {
	Fetch(ctx context.Context, url vo.URL) (*entity.Game8MonsterSourceData, error)
}
