package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/iotassss/puzzdra-monster-rating/internal/domain/service"
	"github.com/iotassss/puzzdra-monster-rating/internal/infrastructure/cli"
	"github.com/iotassss/puzzdra-monster-rating/internal/presenter"
	loader "github.com/iotassss/puzzdra-monster-rating/internal/repository/file_loader"
	repository "github.com/iotassss/puzzdra-monster-rating/internal/repository/gorm"
	scraper "github.com/iotassss/puzzdra-monster-rating/internal/repository/scraper/game8"
	"github.com/iotassss/puzzdra-monster-rating/internal/usecase"
)

const (
	game8MonsterURLOutputFile      = "data/game8_monster_urls"
	makeMonsterRetryCount          = 3
	makeGame8MonsterURLsRetryCount = 3
)

func main() {
	// logger file
	loggerFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		slog.Error("failed to open log file", slog.Any("error", err))
		return
	}
	defer loggerFile.Close()

	handler := slog.NewJSONHandler(loggerFile, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	// app
	app, err := cli.New20241019App(true)
	if err != nil {
		slog.Error("app initialization failed", slog.Any("error", err))
		return
	}

	// repository
	monsterSourceDataRepo := repository.NewMonsterSourceDataRepository(app.DB)
	monsterRepo := repository.NewMonsterRepository(app.DB)
	game8MonsterRepo := repository.NewGame8MonsterRepository(app.DB)

	// file loader
	monsterSourceDataLoader := loader.NewMonsterSourceDataLoader(
		"data/monsters.json",
		true,
		"https://padmdb.rainbowsite.net/listJson/monster_data.json",
		true,
	)
	game8MonsterURLLoader := loader.NewGame8MonsterURLLoader(game8MonsterURLOutputFile)

	// scraper
	scraperConfig := &scraper.Game8MonsterScraperConfig{
		TimeoutSecond: 5,
		WaitSecond:    2,
		UserAgent:     "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36",
		Debug:         true,
	}
	game8MonsterSourceDataFetcher := scraper.NewGame8MonsterScraper(scraperConfig)
	urlScraperConfig := &scraper.Game8MonsterURLScraperConfig{
		TimeoutSecond: 5,
		UserAgent:     "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36",
		OutputFile:    game8MonsterURLOutputFile,
		Debug:         true,
	}
	game8MonsterURLScraper := scraper.NewGame8MonsterURLScraper(urlScraperConfig)

	// service
	findOriginMonsterByNoService := service.NewFindOriginMonsterByNoService(
		monsterRepo,
		monsterSourceDataRepo,
	)
	convertGame8SourceDataService := service.NewConvertGame8SourceDataSV(
		findOriginMonsterByNoService,
	)

	// presenter
	createAllMonsterSourceDataPresenter := presenter.NewCreateAllMonsterSourceDataPresenter()
	createAllMonstersPresenter := presenter.NewCreateAllMonstersPresenter()
	createAllGame8MonstersPresenter := presenter.NewCreateAllGame8MonstersPresenter()

	// usecase
	createAllMonsterSourceData := usecase.NewCreateAllMonsterSourceDataUsecaseInteractor(
		createAllMonsterSourceDataPresenter,
		monsterSourceDataRepo,
		monsterSourceDataLoader,
	)
	createAllMonsters := usecase.NewCreateAllMonstersUsecaseInteractor(
		createAllMonstersPresenter,
		monsterRepo,
		monsterSourceDataRepo,
		findOriginMonsterByNoService,
	)
	createAllGame8Monsters := usecase.NewCreateAllGame8MonstersUsecaseInteractor(
		createAllGame8MonstersPresenter,
		game8MonsterRepo,
		game8MonsterURLLoader,
		game8MonsterSourceDataFetcher,
		convertGame8SourceDataService,
	)

	ctx := context.Background()

	// execute
	for i := 0; i < makeMonsterRetryCount; i++ {
		slog.Info("start createAllMonsterSourceData.Execute...")
		fmt.Println("start createAllMonsterSourceData.Execute...")

		err = createAllMonsterSourceData.Execute(ctx)
		if err != nil {
			slog.Error("createAllMonsterSourceData.Execute failed", slog.Any("error", err))
			if i < makeMonsterRetryCount-1 {
				slog.Info("retry createAllMonsterSourceData.Execute")
				continue
			}
			// TODO: ここでエラーを管理者に通知する

			return
		}

		slog.Info("createAllMonsterSourceData.Execute succeeded!")
		fmt.Println("createAllMonsterSourceData.Execute succeeded!")

		slog.Info("start createAllMonsters.Execute...")
		fmt.Println("start createAllMonsters.Execute...")

		err = createAllMonsters.Execute(ctx)
		if err != nil {
			slog.Error("createAllMonsters.Execute failed", slog.Any("error", err))
			if i < makeMonsterRetryCount-1 {
				slog.Info("retry createAllMonsters.Execute")
				continue
			}
			// TODO: ここでエラーを管理者に通知する

			return
		}

		slog.Info("createAllMonsters.Execute succeeded!")
		fmt.Println("createAllMonsters.Execute succeeded!")

		break
	}

	// _ = createAllMonsterSourceData
	// _ = createAllMonsters

	for i := 0; i < makeGame8MonsterURLsRetryCount; i++ {
		slog.Info("start selectGame8MonsterURLs.Execute...")
		fmt.Println("start selectGame8MonsterURLs.Execute...")

		err = game8MonsterURLScraper.Fetch(ctx)
		if err != nil {
			slog.Error("game8MonsterURLScraper.Fetch failed", slog.Any("error", err))
			if i < makeGame8MonsterURLsRetryCount-1 {
				slog.Info("retry game8MonsterURLScraper.Fetch")
				continue
			}
			// TODO: ここでエラーを管理者に通知する

			return
		}

		slog.Info("game8MonsterURLScraper.Fetch succeeded!")
		fmt.Println("game8MonsterURLScraper.Fetch succeeded!")

		break
	}

	slog.Info("start createAllGame8Monsters.Execute...")
	fmt.Println("start createAllGame8Monsters.Execute...")

	err = createAllGame8Monsters.Execute(ctx)
	if err != nil {
		slog.Error("createAllGame8Monsters.Execute failed", slog.Any("error", err))
		// TODO: ここでエラーを管理者に通知する

		return
	}

	slog.Info("createAllGame8Monsters.Execute succeeded!")
	fmt.Println("createAllGame8Monsters.Execute succeeded!")
}
