package telegram

import (
	"go.uber.org/zap"
)

type tgLoggerAdapter struct{}

func (tgLoggerAdapter) Println(v ...interface{}) {
	zap.S().Infoln(v...)
}
func (tgLoggerAdapter) Printf(format string, v ...interface{}) {
	zap.S().Infof(format, v...)
}

var loggerAdapter = tgLoggerAdapter{}
