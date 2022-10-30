package db

import (
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"notify-api/db/entity"
	"notify-api/utils"
)

var DB *gorm.DB

var RWLock = sync.RWMutex{}

func Init() {

	utils.RequireFile("data/notify.db")
	var err error
	DB, err = gorm.Open(sqlite.Open("data/notify.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		zap.S().Fatalf("Failed to open database: %+v", err)
	}

	err = DB.AutoMigrate(
		&entity.Message{},
		&entity.PushToken{},
	)
	if err != nil {
		zap.S().Fatalf("Failed to migrate database: %+v", err)
	}
}
