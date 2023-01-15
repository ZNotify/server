//go:build !test

package config

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
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
	}

	if path == "ENV" {
		data = []byte(os.Getenv("CONFIG"))
	} else {
		// read config file
		data, err = os.ReadFile(path)
		if err != nil {
			fmt.Println("Failed to read config file.")
			fmt.Printf("%+v\n", err)
			os.Exit(1)
		}
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Failed to parse config file.")
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	Config = config
}
