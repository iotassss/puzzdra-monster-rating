package database

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB(
	debug bool,
) (*gorm.DB, error) {
	dsn := "testuser:testpw@tcp(localhost:3306)/puzzdra_db?charset=utf8mb4&parseTime=True&loc=Local"

	var newLogger logger.Interface
	if !debug {
		newLogger = logger.New(
			logger.Writer(nil), // ログ出力先を指定しない（標準出力に出力される）
			logger.Config{
				SlowThreshold: time.Second,   // スロークエリの閾値（無効にすることも可能）
				LogLevel:      logger.Silent, // Silentモードにして全てのログ出力を抑制
				Colorful:      false,         // カラーログを無効化
			},
		)
		return gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})
	}

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})

}
