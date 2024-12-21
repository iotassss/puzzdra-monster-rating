package usecase

import (
	"context"
	"fmt"

	"github.com/iotassss/puzzdra-monster-rating/internal/domain/repository"
)

type CreateAllMonsterSourceDataPresenter interface {
	Present() error
}

type CreateAllMonsterSourceDataUsecase interface {
	Execute(ctx context.Context) error
}

type CreateAllMonsterSourceDataUsecaseInteractor struct {
	presenter               CreateAllMonsterSourceDataPresenter
	monsterSourceDataRepo   repository.MonsterSourceDataRepository
	monsterSourceDataLoader repository.MonsterSourceDataLoader
}

func NewCreateAllMonsterSourceDataUsecaseInteractor(
	presenter CreateAllMonsterSourceDataPresenter,
	monsterSourceDataRepo repository.MonsterSourceDataRepository,
	monsterSourceDataLoader repository.MonsterSourceDataLoader,
) *CreateAllMonsterSourceDataUsecaseInteractor {
	return &CreateAllMonsterSourceDataUsecaseInteractor{
		presenter:               presenter,
		monsterSourceDataRepo:   monsterSourceDataRepo,
		monsterSourceDataLoader: monsterSourceDataLoader,
	}
}

func (uc *CreateAllMonsterSourceDataUsecaseInteractor) Execute(ctx context.Context) error {
	// *caution: all MonsterSourceData is loaded into memory
	rawMonsterSourceDataList, err := uc.monsterSourceDataLoader.LoadAll(ctx)
	if err != nil {
		return err
	}

	for index, rawMonsterSourceData := range rawMonsterSourceDataList {
		err = uc.monsterSourceDataRepo.Save(ctx, rawMonsterSourceData)
		if err != nil {
			return err
		}

		// debug: 一時的に進捗バーを表示
		displayProgressBar(index+1, len(rawMonsterSourceDataList))
	}

	return uc.presenter.Present()
}

// --- 以下、いつか消す ---

// プログレスバーを表示する関数
func displayProgressBar(current, total int) {
	barLength := 100 // プログレスバーの長さ
	filled := (current * barLength) / total
	empty := barLength - filled

	// プログレスバーの表示
	fmt.Printf("\r%s%s(%d/%d)", repeat("■", filled), repeat("□", empty), current, total)
}

// 繰り返し文字列を生成する関数
func repeat(char string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += char
	}
	return result
}
