package cmd

import (
	"strconv"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"

	"notify-api/setup"
	config2 "notify-api/setup/config"
)

func start(ctx *cli.Context) error {
	host := ctx.String("host")
	port := ctx.Int("port")

	path := ctx.String("config")
	config2.Load(path)

	if host != "" {
		config2.Config.Server.Host = host
	}
	if port != 0 {
		config2.Config.Server.Port = port
	}

	var address string
	if ctx.String("address") != "" {
		address = ctx.String("address")
	} else {
		address = config2.Config.Server.Host + ":" + strconv.Itoa(config2.Config.Server.Port)
	}

	zap.S().Infof("Server is running on %s", address)

	setup.Setup()
	err := setup.NewRouter().Run(address)
	if err != nil {
		return err
	}
	return nil
}
