//go:build test

package config

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Load(path string) {
	zap.S().Infof("Running a test instance, using default configuration.")
	Config = Configuration{
		Server: ServerConfiguration{
			Port: 14444,
			Host: "0.0.0.0",
			Mode: TestMode,
		},
		Database: DatabaseConfiguration{
			Type: Sqlite,
			DSN:  "data/notify.db",
		},
		Senders: make(map[string]SenderConfiguration),
		Users:   make(UserConfiguration, 0),
	}
	Config.Users = append(Config.Users, "test")
	gin.SetMode(gin.TestMode)
}
