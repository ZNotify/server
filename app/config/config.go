//go:build !test

package config

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

func Load(path string) *Configuration {
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
			Address: "0.0.0.0:14444",
			Mode:    mode,
		},
		Database: DatabaseConfiguration{
			Type: Sqlite,
			DSN:  "data/notify.db",
		},
		Senders: SenderConfiguration{
			WebSocket: true,
		},
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
	return &config
}
