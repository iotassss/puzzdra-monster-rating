package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	webhandler "github.com/iotassss/puzzdra-monster-rating/internal/handler/web"
	app "github.com/iotassss/puzzdra-monster-rating/internal/infrastructure/web"
	presenter "github.com/iotassss/puzzdra-monster-rating/internal/presenter/web"
	repository "github.com/iotassss/puzzdra-monster-rating/internal/repository/gorm"
	"github.com/iotassss/puzzdra-monster-rating/internal/usecase"
	"github.com/joho/godotenv"
)

func main() {
	debug := true

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
	app, err := app.New20241030App(debug)
	if err != nil {
		slog.Error("app initialization failed", slog.Any("error", err))
		return
	}

	// load .env
	err = godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", slog.Any("error", err))
	}

	// router
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*.html")

	// usecase
	getMonsterRatingPresenter := presenter.NewGetMonsterRatingPresenter()
	monsterRepo := repository.NewMonsterRepository(app.DB)
	game8MonsterRepo := repository.NewGame8MonsterRepository(app.DB)
	getMonsterRating := usecase.NewGetMonsterRatingUsecaseInteractor(
		getMonsterRatingPresenter,
		monsterRepo,
		game8MonsterRepo,
	)

	// handler
	getMonsterRatingHandler := webhandler.NewGetMonsterRatingHandler(
		app.DB,
		getMonsterRating,
	)

	// get_monster_ratingのエンドポイント
	router.GET("/puzzdra/monster/rating/:no", func(c *gin.Context) {
		getMonsterRatingPresenter.SetGinContext(c)
		getMonsterRatingHandler.Execute(c)
	})

	// サーバーを起動
	// TODO: 環境変数からポート番号を取得する
	if err := router.Run(":8080"); err != nil {
		slog.Error("failed to start server", slog.Any("error", err))
		return
	}
}
