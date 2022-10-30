package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

//goland:noinspection GoUnusedGlobalVariable,SpellCheckingInspection
var (
	Desugar     = logger.Desugar
	Named       = logger.Named
	WithOptions = logger.WithOptions
	With        = logger.With
	Debug       = logger.Debug
	Info        = logger.Info
	Warn        = logger.Warn
	Error       = logger.Error
	DPanic      = logger.DPanic
	Panic       = logger.Panic
	Fatal       = logger.Fatal
	Debugf      = logger.Debugf
	Infof       = logger.Infof
	Warnf       = logger.Warnf
	Errorf      = logger.Errorf
	DPanicf     = logger.DPanicf
	Panicf      = logger.Panicf
	Fatalf      = logger.Fatalf
	Debugw      = logger.Debugw
	Infow       = logger.Infow
	Warnw       = logger.Warnw
	Errorw      = logger.Errorw
	DPanicw     = logger.DPanicw
	Panicw      = logger.Panicw
	Fatalw      = logger.Fatalw
	Debugln     = logger.Debugln
	Infoln      = logger.Infoln
	Warnln      = logger.Warnln
	Errorln     = logger.Errorln
	DPanicln    = logger.DPanicln
	Panicln     = logger.Panicln
	Fatalln     = logger.Fatalln
	Sync        = logger.Sync
)

func init() {
	fmt.Println("log init")
	var cfg zap.Config
	if gin.Mode() == gin.ReleaseMode {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}
	cfg.Encoding = "console"
	l, err := cfg.Build()
	if err != err {
		panic(err)
	}
	logger = l.Sugar()
	fmt.Println("log init done")
}
