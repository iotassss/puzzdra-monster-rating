package service

import (
	"context"

	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/entity"
	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"
	"github.com/iotassss/puzzdra-monster-rating/internal/domain/repository"
)

type FindOriginMonsterByNoService interface {
	Execute(ctx context.Context, no vo.No) (*entity.Monster, error)
}

type FindOriginMonsterByNoSV struct {
	monsterRepository           repository.MonsterRepository
	monsterSourceDataRepository repository.MonsterSourceDataRepository
}

func NewFindOriginMonsterByNoService(
	monsterRepository repository.MonsterRepository,
	monsterSourceDataRepository repository.MonsterSourceDataRepository,
) *FindOriginMonsterByNoSV {
	return &FindOriginMonsterByNoSV{
		monsterRepository:           monsterRepository,
		monsterSourceDataRepository: monsterSourceDataRepository,
	}
}

func (s *FindOriginMonsterByNoSV) Execute(ctx context.Context, no vo.No) (*entity.Monster, error) {
	// 処理内容
	// monsterSourceDataのbaseNoを辿って、originMonsterを取得する
	// originMonsterが見つからない場合は、そのモンスターがその進化系統の起源（origin）。nilを返す

	monsterSourceData, err := s.monsterSourceDataRepository.FindByNo(ctx, no)
	if err != nil {
		return nil, err
	}
	if isOriginMonster(monsterSourceData) {
		return nil, nil
	}

	for monsterSourceData.BaseNo() != nil {
		monsterSourceData, err = s.monsterSourceDataRepository.FindByNo(ctx, *monsterSourceData.BaseNo())
		if err != nil {
			return nil, err
		}
		if isOriginMonster(monsterSourceData) {
			break
		}
	}

	exists, err := s.monsterRepository.Exists(ctx, monsterSourceData.No())
	if err != nil {
		return nil, err
	}

	if exists {
		originMonster, err := s.monsterRepository.FindByNo(ctx, monsterSourceData.No())
		if err != nil {
			return nil, err
		}
		return originMonster, nil
	}

	originMonster := entity.NewMonster(
		vo.NewTemporaryID(),
		monsterSourceData.No(),
		monsterSourceData.Name(),
		nil,
	)

	return originMonster, nil
}

func isOriginMonster(monsterSourceData *entity.MonsterSourceData) bool {
	return monsterSourceData.BaseNo() == nil
}
