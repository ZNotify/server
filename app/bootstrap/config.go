package bootstrap

import (
	"net/url"

	"github.com/ZNotify/server/app/config"
	"github.com/ZNotify/server/app/global"
)

func initializeConfig(args Args) {
	path := args.Config
	if path == "" {
		path = "data/config.yaml"
	}

	c := config.Load(path)

	var address string
	if args.Address != "" {
		address = args.Address
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
