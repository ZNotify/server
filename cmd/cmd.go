package cmd

import (
	"github.com/urfave/cli/v2"
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
	Action: start,
}
