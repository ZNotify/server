package cmd

import (
	"github.com/urfave/cli/v2"

	"github.com/ZNotify/server/cmd/commands"
)

var Cli = &cli.App{
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
			Name:  "address",
			Usage: "Set listen address to `ADDRESS`.",
		},
	},
	Action: commands.Run,
}
