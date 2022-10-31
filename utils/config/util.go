package config

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"notify-api/utils"
	"os"
)

func IsTest() bool {
	return Config.Server.Mode == TestMode
}

func IsDev() bool {
	return Config.Server.Mode == DevMode
}

func IsProd() bool {
	return Config.Server.Mode == ProdMode
}

func Set(c Configuration) {
	Config = c
}

func Load(path string) {
	var data []byte
	var err error

	if path == "ENV" {
		data = []byte(os.Getenv("CONFIG"))
	} else {
		// read config file
		data, err = os.ReadFile(path)
		if err != nil {
			panic(err)
		}
	}

	var mode string
	if utils.IsTestInstance() {
		mode = TestMode
	} else if gin.Mode() == gin.DebugMode {
		mode = DevMode
	} else {
		mode = ProdMode
	}

	config := Configuration{
		Server: ServerConfiguration{
			Port: 14444,
			Host: "127.0.0.1",
			Mode: mode,
		},
		Database: DatabaseConfiguration{
			Type: Sqlite,
			DSN:  "data/notify.db",
		},
		Senders: make(map[string]SenderConfiguration),
		Users:   make(UserConfiguration, 0),
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	Config = config
}
