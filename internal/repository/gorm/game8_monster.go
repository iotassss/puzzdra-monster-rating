package repository

import (
	"context"

	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/entity"
	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"
	"gorm.io/gorm"
)

type Game8MonsterScore struct {
	gorm.Model
	Game8MonsterID uint   `gorm:"not null"`
	Name           string `gorm:"not null"`
	LeaderPoint    string `gorm:"not null"`
	SubLeaderPoint string `gorm:"not null"`
	AssistPoint    string `gorm:"not null"`
}

type Game8Monster struct {
	gorm.Model
	OriginMonsterNo int                  `gorm:"unique;not null"`
	URL             string               `gorm:"unique;not null"`
	Scores          []*Game8MonsterScore `gorm:"foreignKey:Game8MonsterID"`
}

type Game8MonsterRepository struct {
	db *gorm.DB
}

func NewGame8MonsterRepository(db *gorm.DB) *Game8MonsterRepository {
	return &Game8MonsterRepository{db: db}
}

func (r *Game8MonsterRepository) FindByNo(ctx context.Context, no vo.No) (*entity.Game8Monster, error) {
	var gormGame8Monster *Game8Monster
	if err := r.db.WithContext(ctx).
		Where("origin_monster_no = ?", no.Value()).
		First(&gormGame8Monster).Error; err != nil {
		return nil, err
	}

	game8Monster, err := r.convertToEntity(gormGame8Monster)
	if err != nil {
		return nil, err
	}

	return game8Monster, nil
}

func (r *Game8MonsterRepository) Exists(ctx context.Context, no vo.No) (bool, error) {
	var game8Monster *Game8Monster
	if err := r.db.WithContext(ctx).
		Where("origin_monster_no = ?", no.Value()).
		First(&game8Monster).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *Game8MonsterRepository) Save(ctx context.Context, game8Monster *entity.Game8Monster) error {
	var scores []*Game8MonsterScore
	for _, score := range game8Monster.Scores() {
		scores = append(scores, &Game8MonsterScore{
			Name:           score.Name().Value(),
			LeaderPoint:    score.LeaderPoint().Value(),
			SubLeaderPoint: score.SubLeaderPoint().Value(),
			AssistPoint:    score.AssistPoint().Value(),
		})
	}

	url := game8Monster.URL().Value()
	gormGame8Monster := Game8Monster{
		Model: gorm.Model{
			ID: uint(game8Monster.ID().Value()),
		},
		OriginMonsterNo: game8Monster.OriginMonsterNo().Value(),
		URL:             url.String(),
		Scores:          scores,
	}

	if err := r.db.WithContext(ctx).Save(&gormGame8Monster).Error; err != nil {
		return err
	}

	return nil
}

func (r *Game8MonsterRepository) convertToEntity(gormGame8Monster *Game8Monster) (*entity.Game8Monster, error) {
	var scores []*entity.Game8MonsterScore
	for _, gormScore := range gormGame8Monster.Scores {
		id, err := vo.NewID(int(gormScore.ID))
		if err != nil {
			return nil, err
		}
		name, err := vo.NewMonsterName(gormScore.Name)
		if err != nil {
			return nil, err
		}
		leaderPoint, err := vo.NewGame8MonsterPoint(gormScore.LeaderPoint)
		if err != nil {
			return nil, err
		}
		subLeaderPoint, err := vo.NewGame8MonsterPoint(gormScore.SubLeaderPoint)
		if err != nil {
			return nil, err
		}
		assistPoint, err := vo.NewGame8MonsterPoint(gormScore.AssistPoint)
		if err != nil {
			return nil, err
		}

		scores = append(scores, entity.NewGame8MonsterScore(
			id,
			name,
			leaderPoint,
			subLeaderPoint,
			assistPoint,
		))
	}

	id, err := vo.NewID(int(gormGame8Monster.ID))
	if err != nil {
		return nil, err
	}
	no, err := vo.NewNo(gormGame8Monster.OriginMonsterNo)
	if err != nil {
		return nil, err
	}
	url, err := vo.NewURL(gormGame8Monster.URL)
	if err != nil {
		return nil, err
	}

	return entity.NewGame8Monster(
		id,
		no,
		url,
		scores,
	), nil
}