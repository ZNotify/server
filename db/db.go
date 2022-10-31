package db

import (
	"sync"

	"notify-api/utils/config"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"notify-api/db/entity"
	"notify-api/utils"
)

var DB *gorm.DB

var RWLock = sync.RWMutex{}

func setupSqlite() gorm.Dialector {
	path := config.Config.Database.DSN
	utils.RequireFile(path)
	return sqlite.Open(path)
}

func setupMysql() gorm.Dialector {
	return mysql.Open(config.Config.Database.DSN)
}

func setupPgsql() gorm.Dialector {
	return postgres.Open(config.Config.Database.DSN)
}

func dbConnection() gorm.Dialector {
	switch config.Config.Database.Type {
	case config.Sqlite:
		return setupSqlite()
	case config.Mysql:
		return setupMysql()
	case config.Pgsql:
		return setupPgsql()
	default:
		panic("Unsupported database type")
	}
}

func Init() {
	var err error

	DB, err = gorm.Open(dbConnection(), &gorm.Config{
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
