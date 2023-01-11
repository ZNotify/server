package misc

import (
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"notify-api/setup/config"
)

func RequireNetwork() {
	if !config.IsProd() {
		zap.S().Info("Skip connection check in non-production mode")
		return
	}

	go func() {
		_, err := http.Get("https://www.google.com/robots.txt")
		if err != nil {
			zap.S().Panicf("Failed to connect to internet: %v", err)
		}
	}()
}

func RequireX64() {
	if strconv.IntSize != 64 {
		zap.S().Panic("Only 64-bit platform is supported")
	}
}
