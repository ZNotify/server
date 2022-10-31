package cmd

import (
	"github.com/urfave/cli/v2"
	"notify-api/utils/config"
	"notify-api/utils/setup"
	"strconv"
)

var App = &cli.App{
	Name:  "Notify API",
	Usage: "This is Znotify api server.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Load configuration from `FILE`, or use ENV to load from environment variable CONFIG.",
			Value:   "data/config.yaml",
		},
	},
	Action: func(ctx *cli.Context) error {
		path := ctx.String("config")
		config.Load(path)
		address := config.Config.Server.Host + ":" + strconv.Itoa(config.Config.Server.Port)
		err := setup.New().Run(address)
		return err
	},
}
