package bootstrap

import (
	"context"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"

	"notify-api/config"
	"notify-api/db/ent/generate"
	"notify-api/global"
)

func getDatabase() *generate.Client {
	var err error

	var db string
	switch global.App.Config.Database.Type {
	case config.Sqlite:
		db = dialect.SQLite
	case config.Mysql:
		db = dialect.MySQL
	case config.Pgsql:
		db = dialect.Postgres
	default:
		panic("Unsupported database type")
	}

	client, err := generate.Open(db, global.App.Config.Database.DSN)
	if err != nil {
		zap.S().Fatalf("Failed to connect database: %+v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		zap.S().Fatalf("Failed to create schema resources: %+v", err)
	}
	return client
}

func initializeDatabase() {
	global.App.DB = getDatabase()
}
