package repository

import (
	"context"
	"errors"
	"log/slog"

	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/entity"
	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"
	"github.com/iotassss/puzzdra-monster-rating/internal/infrastructure/stacktrace"
	"gorm.io/gorm"
)

type Monster struct {
	gorm.Model
	No              int    `gorm:"unique;not null"`
	Name            string `gorm:"size:255;not null"`
	OriginMonsterID *int
	OriginMonster   *Monster `gorm:"foreignKey:OriginMonsterID"`
}

type MonsterRepository struct {
	db *gorm.DB
}

func NewMonsterRepository(db *gorm.DB) *MonsterRepository {
	return &MonsterRepository{db: db}
}

func (r *MonsterRepository) FindByNo(ctx context.Context, no vo.No) (*entity.Monster, error) {
	var gormMonster Monster
	if err := r.db.WithContext(ctx).
		Where("no = ?", no.Value()).
		First(&gormMonster).Error; err != nil {
		slog.Error("failed to find monster by no", slog.Any("error", err), slog.String("stacktrace", stacktrace.Print()))
		return nil, err
	}

	monster, err := r.convertToEntity(&gormMonster)
	if err != nil {
		slog.Error("failed to convert gorm monster to entity", slog.Any("error", err), slog.String("stacktrace", stacktrace.Print()))
		return nil, err
	}

	return monster, nil
}

func (r *MonsterRepository) Exists(ctx context.Context, no vo.No) (bool, error) {
	var monster *Monster
	if err := r.db.WithContext(ctx).
		Where("no = ?", no.Value()).
		First(&monster).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		slog.Error("failed to find monster by no", slog.Any("error", err), slog.String("stacktrace", stacktrace.Print()))
		return false, err
	}

	return true, nil
}

func (r *MonsterRepository) Save(ctx context.Context, monster *entity.Monster) error {
	var originMonsterID *int
	if monster.OriginMonster() != nil {
		originMonsterID = new(int)
		*originMonsterID = int(monster.OriginMonster().ID().Value())
	}

	exists, err := r.Exists(ctx, monster.No())
	if err != nil {
		return err
	}

	if exists {
		var gormMonster Monster
		if err := r.db.WithContext(ctx).
			Where("no = ?", monster.No().Value()).
			First(&gormMonster).Error; err != nil {
			slog.Error("failed to find monster by no", slog.Any("error", err), slog.String("stacktrace", stacktrace.Print()))
			return err
		}

		gormMonster.Name = monster.Name().Value()
		gormMonster.OriginMonsterID = originMonsterID

		if err := r.db.WithContext(ctx).Save(&gormMonster).Error; err != nil {
			slog.Error("failed to save monster", slog.Any("error", err), slog.String("stacktrace", stacktrace.Print()))
			return err
		}
	} else {
		gormMonster := Monster{
			No:              monster.No().Value(),
			Name:            monster.Name().Value(),
			OriginMonsterID: originMonsterID,
		}

		if err := r.db.WithContext(ctx).Create(&gormMonster).Error; err != nil {
			slog.Error("failed to create monster", slog.Any("error", err), slog.String("stacktrace", stacktrace.Print()))
			return err
		}

		assignedID, err := vo.NewID(int(gormMonster.ID))
		if err != nil {
			slog.Error("failed to create ID", slog.Any("error", err), slog.String("stacktrace", stacktrace.Print()))
			return err
		}
		monster.SetID(assignedID)
	}

	return nil
}

func (r *MonsterRepository) convertToEntity(gormMonster *Monster) (*entity.Monster, error) {
	monster, err := r.convertToEntityWithoutOriginMonster(gormMonster)
	if err != nil {
		return nil, err
	}

	if gormMonster.OriginMonster != nil {
		originMonster, err := r.convertToEntityWithoutOriginMonster(gormMonster.OriginMonster)
		if err != nil {
			return nil, err
		}
		monster.SetOriginMonster(originMonster)
	}

	return monster, nil
}

func (r *MonsterRepository) convertToEntityWithoutOriginMonster(gormMonster *Monster) (*entity.Monster, error) {
	if gormMonster == nil {
		return nil, errors.New("gormMonster is nil")
	}

	id, err := vo.NewID(int(gormMonster.ID))
	if err != nil {
		return nil, err
	}
	no, err := vo.NewNo(gormMonster.No)
	if err != nil {
		return nil, err
	}
	name, err := vo.NewMonsterName(gormMonster.Name)
	if err != nil {
		return nil, err
	}

	monster := entity.NewMonster(id, no, name, nil)

	return monster, nil
}
