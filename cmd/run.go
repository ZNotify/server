package cmd

import (
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"

	"notify-api/app/router"
	"notify-api/bootstrap"
	"notify-api/global"
)

func run(ctx *cli.Context) error {
	bootstrap.BootStrap(ctx)
	r := router.NewRouter()
	err := r.Run(global.App.Config.Server.Address)
	if err != nil {
		zap.S().Fatalf("Failed to start server: %+v", err)
	}
	defer func() {
		err = global.App.DB.Close()
		if err != nil {
			zap.S().Errorf("Failed to close database: %+v", err)
		}
		err = zap.L().Sync()
		if err != nil {
			zap.S().Errorf("Failed to sync logger: %+v", err)
		}
	}()
	return nil
}
