package bootstrap

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/ZNotify/server/app/utils"
)

func initializeLog() {
	var pe zapcore.EncoderConfig
	if gin.Mode() == gin.ReleaseMode {
		pe = zap.NewProductionEncoderConfig()
	} else {
		pe = zap.NewDevelopmentEncoderConfig()
	}

	utils.RequireFile("data/app.log")
	logFile, err := os.OpenFile("data/app.log", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	fileEncoder := zapcore.NewJSONEncoder(pe)

	pe.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(pe)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(logFile), zapcore.DebugLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.InfoLevel),
	)

	l := zap.New(core, zap.AddStacktrace(zapcore.WarnLevel))
	if err != err {
		panic(err)
	}
	zap.ReplaceGlobals(l)
}
