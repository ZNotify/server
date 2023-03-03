package bootstrap

import (
	"github.com/ZNotify/server/app/config"
	"github.com/ZNotify/server/app/global"
)

func initializeConfig(args Args) {
	cfg := MergeConfig(args)
	err, errS := cfg.Validate()
	if err != nil {
		panic(errS)
	}

	global.App.Config = cfg
}

func MergeConfig(args Args) *config.Configuration {
	path := args.ConfigPath
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

	return c
}
