package usecase_test

import (
	"context"
	"testing"

	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/entity"
	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"
	"github.com/iotassss/puzzdra-monster-rating/internal/usecase"
)

// stub presenter
type stubCreateAllMonsterSourceDataPresenter struct{}

func (p *stubCreateAllMonsterSourceDataPresenter) Present() error {
	return nil
}

type stubMonsterSourceDataRepository struct {
	FindAllCallCount  int
	MonsterSourceData []*entity.MonsterSourceData
	FindByNoCallCount int
	FindByNoReturn    []*entity.MonsterSourceData
	ExistsCallCount   int
	ExistsReturn      []bool
	SaveCallCount     int
	SaveError         []error
}

func (r *stubMonsterSourceDataRepository) FindAll(
	ctx context.Context,
) ([]*entity.MonsterSourceData, error) {
	return r.MonsterSourceData, nil
}

func (r *stubMonsterSourceDataRepository) FindByNo(
	ctx context.Context,
	no vo.No,
) (*entity.MonsterSourceData, error) {
	monsterSourceData := r.MonsterSourceData[r.FindByNoCallCount]
	r.FindByNoCallCount++
	return monsterSourceData, nil
}

func (r *stubMonsterSourceDataRepository) Exists(ctx context.Context, no vo.No) (bool, error) {
	exists := r.ExistsReturn[r.ExistsCallCount]
	r.ExistsCallCount++
	return exists, nil
}

func (r *stubMonsterSourceDataRepository) Save(
	ctx context.Context,
	monsterSourceData *entity.MonsterSourceData,
) error {
	err := r.SaveError[r.SaveCallCount]
	r.SaveCallCount++
	return err
}

type stubMonsterSourceDataLoader struct {
	MonsterSourceData []*entity.MonsterSourceData
}

func (l *stubMonsterSourceDataLoader) LoadAll(
	ctx context.Context,
) ([]*entity.MonsterSourceData, error) {
	return l.MonsterSourceData, nil
}

func TestCreateAllMonsterSourceDataUsecaseInteractor_Execute(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		presenter := &stubCreateAllMonsterSourceDataPresenter{}

		monsterSourceDataRepo := &stubMonsterSourceDataRepository{
			SaveError: []error{nil, nil, nil},
		}

		monsterSourceDataID1, err := vo.NewID(1)
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		monsterSourceDataNo1, err := vo.NewNo(1)
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		name1, err := vo.NewMonsterName("name1")
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		baseNo1, err := vo.NewNo(11)
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		monsterSourceData1 := entity.NewMonsterSourceData(
			monsterSourceDataID1,
			monsterSourceDataNo1,
			name1,
			&baseNo1,
		)

		monsterSourceDataID2, err := vo.NewID(2)
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		monsterSourceDataNo2, err := vo.NewNo(2)
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		name2, err := vo.NewMonsterName("name2")
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		baseNo2, err := vo.NewNo(22)
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		monsterSourceData2 := entity.NewMonsterSourceData(
			monsterSourceDataID2,
			monsterSourceDataNo2,
			name2,
			&baseNo2,
		)

		monsterSourceDataID3, err := vo.NewID(3)
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		monsterSourceDataNo3, err := vo.NewNo(3)
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		name3, err := vo.NewMonsterName("name3")
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		baseNo3, err := vo.NewNo(33)
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		monsterSourceData3 := entity.NewMonsterSourceData(
			monsterSourceDataID3,
			monsterSourceDataNo3,
			name3,
			&baseNo3,
		)

		monsterSourceDataLoader := &stubMonsterSourceDataLoader{
			MonsterSourceData: []*entity.MonsterSourceData{
				monsterSourceData1,
				monsterSourceData2,
				monsterSourceData3,
			},
		}

		uc := usecase.NewCreateAllMonsterSourceDataUsecaseInteractor(
			presenter,
			monsterSourceDataRepo,
			monsterSourceDataLoader,
		)

		// execute
		err = uc.Execute(context.Background())

		// verify
		if err != nil {
			t.Errorf("error should be nil, but got: %v", err)
		}
	})
}
