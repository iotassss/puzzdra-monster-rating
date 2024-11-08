package usecase_test

import (
	"context"
	"testing"

	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/entity"
	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"
	"github.com/iotassss/puzzdra-monster-rating/internal/usecase"
)

// stub presenter
type stubCreateAllMonstersPresenter struct{}

func (p *stubCreateAllMonstersPresenter) Present() error {
	return nil
}

type stubMonsterRepository struct {
	FindByNoCallCount int
	Monsters          []*entity.Monster
	ExistsCallCount   int
	ExistsReturn      []bool
	SaveCallCount     int
	SaveError         []error
}

func (r *stubMonsterRepository) FindByNo(
	ctx context.Context,
	no vo.No,
) (*entity.Monster, error) {
	monster := r.Monsters[r.FindByNoCallCount]
	r.FindByNoCallCount++
	return monster, nil
}

func (r *stubMonsterRepository) Exists(
	ctx context.Context,
	no vo.No,
) (bool, error) {
	exists := r.ExistsReturn[r.ExistsCallCount]
	r.ExistsCallCount++
	return exists, nil
}

func (r *stubMonsterRepository) Save(
	ctx context.Context,
	monster *entity.Monster,
) error {
	err := r.SaveError[r.SaveCallCount]
	r.SaveCallCount++
	return err
}

type stubFindOriginMonsterByNoService struct {
	CallCount     int
	originMonster []*entity.Monster
}

func (s *stubFindOriginMonsterByNoService) Execute(
	ctx context.Context,
	no vo.No,
) (*entity.Monster, error) {
	originMonster := s.originMonster[s.CallCount]
	s.CallCount++
	return originMonster, nil
}

func TestCreateAllMonstersUsecaseInteractor_Execute(t *testing.T) {
	t.Run("success", func(t *testing.T) {

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

		stubMonsterSourceDataRepo := &stubMonsterSourceDataRepository{
			MonsterSourceData: []*entity.MonsterSourceData{
				monsterSourceData1,
				monsterSourceData2,
				monsterSourceData3,
			},
		}

		monsterID1, err := vo.NewID(1)
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		monsterNo1, err := vo.NewNo(1)
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		monsterName1, err := vo.NewMonsterName("name1")
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		originMonsterID1, err := vo.NewID(1)
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		originMonsterNo1, err := vo.NewNo(11)
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		originMonsterName1, err := vo.NewMonsterName("originName1")
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		originMonster1 := entity.NewMonster(
			originMonsterID1,
			originMonsterNo1,
			originMonsterName1,
			nil,
		)
		monster1 := entity.NewMonster(
			monsterID1,
			monsterNo1,
			monsterName1,
			originMonster1,
		)

		monsterID2, err := vo.NewID(2)
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		monsterNo2, err := vo.NewNo(2)
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		monsterName2, err := vo.NewMonsterName("name2")
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		originMonsterID2, err := vo.NewID(2)
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		originMonsterNo2, err := vo.NewNo(22)
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		originMonsterName2, err := vo.NewMonsterName("originName2")
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		originMonster2 := entity.NewMonster(
			originMonsterID2,
			originMonsterNo2,
			originMonsterName2,
			nil,
		)
		monster2 := entity.NewMonster(
			monsterID2,
			monsterNo2,
			monsterName2,
			originMonster2,
		)

		monsterID3, err := vo.NewID(3)
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		monsterNo3, err := vo.NewNo(3)
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		monsterName3, err := vo.NewMonsterName("name3")
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		originMonsterID3, err := vo.NewID(3)
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		originMonsterNo3, err := vo.NewNo(33)
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		originMonsterName3, err := vo.NewMonsterName("originName3")
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
		originMonster3 := entity.NewMonster(
			originMonsterID3,
			originMonsterNo3,
			originMonsterName3,
			nil,
		)
		monster3 := entity.NewMonster(
			monsterID3,
			monsterNo3,
			monsterName3,
			originMonster3,
		)

		stubMonsterRepo := &stubMonsterRepository{
			Monsters: []*entity.Monster{
				monster1,
				monster2,
				monster3,
			},
			ExistsReturn: []bool{
				true,
				false,
				false,
			},
			SaveError: []error{
				nil,
				nil,
				nil,
			},
		}

		stubFindOriginMonsterByNoService := &stubFindOriginMonsterByNoService{
			originMonster: []*entity.Monster{
				originMonster1,
				originMonster2,
				originMonster3,
			},
		}

		uc := usecase.NewCreateAllMonstersUsecaseInteractor(
			&stubCreateAllMonstersPresenter{},
			stubMonsterRepo,
			stubMonsterSourceDataRepo,
			stubFindOriginMonsterByNoService,
		)

		err = uc.Execute(context.Background())
		if err != nil {
			t.Fatalf("error should be nil, but got: %v", err)
		}
	})
}
