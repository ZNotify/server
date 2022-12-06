//go:build !test

package config

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"os"
)

func Load(path string) {
	var data []byte
	var err error
	var mode string
	if gin.Mode() == gin.DebugMode {
		mode = DevMode
	} else {
		mode = ProdMode
	}

	config := Configuration{
		Server: ServerConfiguration{
			Port: 14444,
			Host: "0.0.0.0",
			Mode: mode,
		},
		Database: DatabaseConfiguration{
			Type: Sqlite,
			DSN:  "data/notify.db",
		},
		Senders: make(map[string]SenderConfiguration),
		Users:   make(UserConfiguration, 0),
	}

	if mode == TestMode {
		Config = config
		return
	}

	if path == "ENV" {
		data = []byte(os.Getenv("CONFIG"))
	} else {
		// read config file
		data, err = os.ReadFile(path)
		if err != nil {
			panic(err)
		}
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	Config = config
}
