package cmd

import (
	"strconv"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"

	"notify-api/utils/config"
	"notify-api/utils/setup"
)

var App = &cli.App{
	Name:                 "Notify API",
	Usage:                "This is Znotify api server.",
	EnableBashCompletion: true,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Load configuration from `FILE`, or use ENV to load from environment variable CONFIG.",
			Value:   "data/config.yaml",
		},
		&cli.StringFlag{
			Name:  "host",
			Usage: "Set listen host to `HOST`.",
		},
		&cli.IntFlag{
			Name:  "port",
			Usage: "Set listen port to `PORT`.",
		},
		&cli.StringFlag{
			Name:  "address",
			Usage: "Set listen address to `ADDRESS`.",
		},
	},
	Action: func(ctx *cli.Context) error {
		host := ctx.String("host")
		port := ctx.Int("port")

		path := ctx.String("config")
		config.Load(path)

		if host != "" {
			config.Config.Server.Host = host
		}
		if port != 0 {
			config.Config.Server.Port = port
		}

		var address string
		if ctx.String("address") != "" {
			address = ctx.String("address")
		} else {
			address = config.Config.Server.Host + ":" + strconv.Itoa(config.Config.Server.Port)
		}

		err := setup.New().Run(address)
		if err != nil {
			return err
		}
		zap.S().Infof("Server is running on %s", address)
		return nil
	},
}
