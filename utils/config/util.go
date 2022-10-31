package config

import (
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"

	"notify-api/utils"
)

func IsTest() bool {
	return Config.Server.Mode == TestMode
}

func SetTest() {
	Config.Server.Mode = TestMode
}

func IsDev() bool {
	return Config.Server.Mode == DevMode
}

func IsProd() bool {
	return Config.Server.Mode == ProdMode
}

func Load(path string) {
	var data []byte
	var err error
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
