package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/ZNotify/server/app/api/router"
	"github.com/ZNotify/server/app/bootstrap"
	"github.com/ZNotify/server/app/global"
)

func Run(cmd *cobra.Command, args []string) {
	bootstrap.BootStrap(bootstrap.Args{
		ConfigPath: cmd.Flag("config").Value.String(),
		Address:    cmd.Flag("address").Value.String(),
	})
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
}
