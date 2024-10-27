package cli

import (
	"log/slog"
	"os"

	database "github.com/iotassss/puzzdra-monster-rating/internal/infrastructure/database/gorm"
	repository "github.com/iotassss/puzzdra-monster-rating/internal/repository/gorm"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type App struct {
	DB     *gorm.DB
	Logger *slog.Logger
	Debug  bool
}

func NewApp(db *gorm.DB) *App {
	return &App{DB: db}
}

func New20241019App(
	loggerFile *os.File,
	debug bool,
) (*App, error) {
	// logger
	handler := slog.NewJSONHandler(loggerFile, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	// load .env
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", slog.Any("error", err))
	}

	// MonsterDataJSONURL := os.Getenv("MONSTER_DATA_JSON_URL")
	// Game8MonsterURLsFile := os.Getenv("GAME8_MONSTER_URLS_FILE")

	// database
	db, err := database.NewDB(debug)
	if err != nil {
		slog.Error("database connection failed", slog.Any("error", err))
		return nil, err
	}

	err = db.AutoMigrate(
		&repository.Monster{},
		&repository.Game8Monster{},
		&repository.Game8MonsterScore{},
		&repository.MonsterSourceData{},
	)
	if err != nil {
		slog.Error("failed to migrate database", slog.Any("error", err))
		return nil, err
	}

	return &App{
		DB:     db,
		Logger: logger,
	}, nil
}
