package bootstrap

import (
	"net/url"

	"github.com/urfave/cli/v2"

	"notify-api/app/config"
	"notify-api/app/global"
)

func initializeConfig(ctx *cli.Context) {
	path := ctx.String("config")

	c := config.Load(path)

	var address string
	if ctx.String("address") != "" {
		address = ctx.String("address")
	} else {
		address = c.Server.Address
	}
	if address == "" {
		address = "0.0.0.0:14444"
	}

	c.Server.Address = address

	// test url
	_, err := url.Parse(c.Server.URL)
	if err != nil {
		panic(err)
	}

	global.App.Config = c
}
