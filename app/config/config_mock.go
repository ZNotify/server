//go:build test

package config

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Load(path string) *Configuration {
	zap.S().Infof("Running a test instance, using default configuration.")
	gin.SetMode(gin.TestMode)
	return &Configuration{
		Server: ServerConfiguration{
			Address: "0.0.0.0:14444",
			Mode:    TestMode,
			URL:     "http://127.0.0.1:14444",
		},
		Database: DatabaseConfiguration{
			Type: Sqlite,
			DSN:  "file:memory:main?mode=memory&cache=shared&_fk=1&_timeout=5000",
		},
		Senders: SenderConfiguration{
			WebSocket: nil,
			FCM:       nil,
			Telegram:  nil,
			WebPush:   nil,
			WNS:       nil,
		},
	}
}
