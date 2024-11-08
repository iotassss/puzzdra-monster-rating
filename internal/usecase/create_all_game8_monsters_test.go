package usecase_test

import (
	"context"
	"testing"

	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/entity"
	"github.com/iotassss/puzzdra-monster-rating/internal/domain/model/vo"
	"github.com/iotassss/puzzdra-monster-rating/internal/usecase"
)

// stub presenter
type stubCreateAllGame8MonstersPresenter struct{}

func (p *stubCreateAllGame8MonstersPresenter) Present() error {
	return nil
}

type stubGame8MonsterRepository struct {
	FindByNoCallCount int
	Game8Monsters     []*entity.Game8Monster
	ExistsCallCount   int
	ExistsReturn      []bool
	SaveCallCount     int
	SaveError         []error
}

func (r *stubGame8MonsterRepository) FindByNo(
	ctx context.Context,
	no vo.No,
) (*entity.Game8Monster, error) {
	game8monster := r.Game8Monsters[r.FindByNoCallCount]
	r.FindByNoCallCount++
	return game8monster, nil
}

func (r *stubGame8MonsterRepository) Exists(
	ctx context.Context,
	no vo.No,
) (bool, error) {
	exists := r.ExistsReturn[r.ExistsCallCount]
	r.ExistsCallCount++
	return exists, nil
}

func (r *stubGame8MonsterRepository) Save(
	ctx context.Context,
	game8Monster *entity.Game8Monster,
) error {
	err := r.SaveError[r.SaveCallCount]
	r.SaveCallCount++
	return err
}

type stubGame8MonsterURLLoader struct {
	URLs []vo.URL
}

func (l *stubGame8MonsterURLLoader) LoadAll(ctx context.Context) ([]vo.URL, error) {
	return l.URLs, nil
}

type stubGame8MonsterSourceDataFetcher struct {
	Game8MonsterSourceData map[string]*entity.Game8MonsterSourceData
}

func (f *stubGame8MonsterSourceDataFetcher) Fetch(
	ctx context.Context,
	url vo.URL,
) (*entity.Game8MonsterSourceData, error) {
	return f.Game8MonsterSourceData[url.Value().String()], nil
}

type stubConvertGame8SourceDataService struct {
	CallCount     int
	Game8Monsters []*entity.Game8Monster
}

func (s *stubConvertGame8SourceDataService) Execute(
	ctx context.Context,
	game8MonsterSourceData *entity.Game8MonsterSourceData,
) (*entity.Game8Monster, error) {
	game8monster := s.Game8Monsters[s.CallCount]
	s.CallCount++
	return game8monster, nil
}

func TestCreateAllGame8MonstersUsecaseInteractor_Execute(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		presenter := &stubCreateAllGame8MonstersPresenter{}

		game8MonsterRepo := &stubGame8MonsterRepository{
			Game8Monsters: []*entity.Game8Monster{{}},
			ExistsReturn:  []bool{false, false, false},
			SaveError:     []error{nil, nil, nil},
		}

		url1, err := vo.NewURL("http://example.com/1")
		if err != nil {
			t.Fatalf("expected err to be nil, but got %v", err)
		}
		url2, err := vo.NewURL("http://example.com/2")
		if err != nil {
			t.Fatalf("expected err to be nil, but got %v", err)
		}
		url3, err := vo.NewURL("http://example.com/3")
		if err != nil {
			t.Fatalf("expected err to be nil, but got %v", err)
		}
		game8MonsterURLLoader := &stubGame8MonsterURLLoader{
			URLs: []vo.URL{url1, url2, url3},
		}

		name1, err := vo.NewMonsterName("name1")
		if err != nil {
			t.Fatalf("expected err to be nil, but got %v", err)
		}
		leaderPoint1, err := vo.NewGame8MonsterPoint("1.0")
		if err != nil {
			t.Fatalf("expected err to be nil, but got %v", err)
		}
		sub1Point1, err := vo.NewGame8MonsterPoint("9.9")
		if err != nil {
			t.Fatalf("expected err to be nil, but got %v", err)
		}
		assistPoint1, err := vo.NewGame8MonsterPoint("-")
		if err != nil {
			t.Fatalf("expected err to be nil, but got %v", err)
		}
		Game8MonsterScoreSourceData1 := entity.NewGame8MonsterScoreSourceData(
			name1,
			leaderPoint1,
			sub1Point1,
			assistPoint1,
		)

		name2, err := vo.NewMonsterName("name2")
		if err != nil {
			t.Fatalf("expected err to be nil, but got %v", err)
		}
		leaderPoint2, err := vo.NewGame8MonsterPoint("2.0")
		if err != nil {
			t.Fatalf("expected err to be nil, but got %v", err)
		}
		sub1Point2, err := vo.NewGame8MonsterPoint("8.8")
		if err != nil {
			t.Fatalf("expected err to be nil, but got %v", err)
		}
		assistPoint2, err := vo.NewGame8MonsterPoint("7.7")
		if err != nil {
			t.Fatalf("expected err to be nil, but got %v", err)
		}
		Game8MonsterScoreSourceData2 := entity.NewGame8MonsterScoreSourceData(
			name2,
			leaderPoint2,
			sub1Point2,
			assistPoint2,
		)

		name3, err := vo.NewMonsterName("name3")
		if err != nil {
			t.Fatalf("expected err to be nil, but got %v", err)
		}
		leaderPoint3, err := vo.NewGame8MonsterPoint("3.0")
		if err != nil {
			t.Fatalf("expected err to be nil, but got %v", err)
		}
		sub1Point3, err := vo.NewGame8MonsterPoint("7.7")
		if err != nil {
			t.Fatalf("expected err to be nil, but got %v", err)
		}
		assistPoint3, err := vo.NewGame8MonsterPoint("6.6")
		if err != nil {
			t.Fatalf("expected err to be nil, but got %v", err)
		}
		Game8MonsterScoreSourceData3 := entity.NewGame8MonsterScoreSourceData(
			name3,
			leaderPoint3,
			sub1Point3,
			assistPoint3,
		)

		no1, err := vo.NewNo(1)
		if err != nil {
			t.Fatalf("expected err to be nil, but got %v", err)
		}
		no2, err := vo.NewNo(2)
		if err != nil {
			t.Fatalf("expected err to be nil, but got %v", err)
		}
		no3, err := vo.NewNo(3)
		if err != nil {
			t.Fatalf("expected err to be nil, but got %v", err)
		}

		game8MonsterSourceData1 := entity.NewGame8MonsterSourceData(
			no1,
			url1,
			[]*entity.Game8MonsterScoreSourceData{Game8MonsterScoreSourceData1},
		)
		game8MonsterSourceData2 := entity.NewGame8MonsterSourceData(
			no2,
			url2,
			[]*entity.Game8MonsterScoreSourceData{Game8MonsterScoreSourceData2},
		)
		game8MonsterSourceData3 := entity.NewGame8MonsterSourceData(
			no3,
			url3,
			[]*entity.Game8MonsterScoreSourceData{Game8MonsterScoreSourceData3},
		)

		game8MonsterSourceDataFetcher := &stubGame8MonsterSourceDataFetcher{
			Game8MonsterSourceData: map[string]*entity.Game8MonsterSourceData{
				url1.Value().String(): game8MonsterSourceData1,
				url2.Value().String(): game8MonsterSourceData2,
				url3.Value().String(): game8MonsterSourceData3,
			},
		}

		scoreID1, err := vo.NewID(1)
		if err != nil {
			t.Fatalf("expected err to be nil, but got %v", err)
		}
		scoreID2, err := vo.NewID(2)
		if err != nil {
			t.Fatalf("expected err to be nil, but got %v", err)
		}
		scoreID3, err := vo.NewID(3)
		if err != nil {
			t.Fatalf("expected err to be nil, but got %v", err)
		}
		game8MonsterScore1 := entity.NewGame8MonsterScore(
			scoreID1,
			name1,
			leaderPoint1,
			sub1Point1,
			assistPoint1,
		)
		game8MonsterScore2 := entity.NewGame8MonsterScore(
			scoreID2,
			name2,
			leaderPoint2,
			sub1Point2,
			assistPoint2,
		)
		game8MonsterScore3 := entity.NewGame8MonsterScore(
			scoreID3,
			name3,
			leaderPoint3,
			sub1Point3,
			assistPoint3,
		)

		game8MonsterID1, err := vo.NewID(1)
		if err != nil {
			t.Fatalf("expected err to be nil, but got %v", err)
		}
		game8MonsterID2, err := vo.NewID(2)
		if err != nil {
			t.Fatalf("expected err to be nil, but got %v", err)
		}
		game8MonsterID3, err := vo.NewID(3)
		if err != nil {
			t.Fatalf("expected err to be nil, but got %v", err)
		}

		game8Monster1 := entity.NewGame8Monster(
			game8MonsterID1,
			no1,
			url1,
			[]*entity.Game8MonsterScore{game8MonsterScore1},
		)
		game8Monster2 := entity.NewGame8Monster(
			game8MonsterID2,
			no2,
			url2,
			[]*entity.Game8MonsterScore{game8MonsterScore2},
		)
		game8Monster3 := entity.NewGame8Monster(
			game8MonsterID3,
			no3,
			url3,
			[]*entity.Game8MonsterScore{game8MonsterScore3},
		)

		convertGame8SourceDataService := &stubConvertGame8SourceDataService{
			Game8Monsters: []*entity.Game8Monster{
				game8Monster1,
				game8Monster2,
				game8Monster3,
			},
		}

		uc := usecase.NewCreateAllGame8MonstersUsecaseInteractor(
			presenter,
			game8MonsterRepo,
			game8MonsterURLLoader,
			game8MonsterSourceDataFetcher,
			convertGame8SourceDataService,
		)
		err = uc.Execute(context.Background())
		if err != nil {
			t.Errorf("expected err to be nil, but got %v", err)
		}
	})
}
