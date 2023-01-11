package db

import (
	"context"

	"entgo.io/ent/dialect"

	"notify-api/ent/generate"
	"notify-api/setup/config"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var C *generate.Client

func Init() {
	var err error

	var db string
	switch config.Config.Database.Type {
	case config.Sqlite:
		db = dialect.SQLite
	case config.Mysql:
		db = dialect.MySQL
	case config.Pgsql:
		db = dialect.Postgres
	default:
		panic("Unsupported database type")
	}

	C, err = generate.Open(db, config.Config.Database.DSN)
	if err != nil {
		zap.S().Fatalf("Failed to connect database: %+v", err)
	}

	if err := C.Schema.Create(context.Background()); err != nil {
		zap.S().Fatalf("Failed to create schema resources: %+v", err)
	}
}
