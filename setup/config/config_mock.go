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
			Host: "127.0.0.1",
			Mode: TestMode,
		},
		Database: DatabaseConfiguration{
			Type: Sqlite,
			DSN:  "file:memory:main?mode=memory&cache=shared&_fk=1",
		},
		Senders: make(map[string]SenderConfiguration),
	}
	gin.SetMode(gin.TestMode)
}