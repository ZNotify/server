//go:build !test

package config

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/invopop/yaml"
)

func Load(path string) *Configuration {
	var data []byte
	var err error
	var mode Mode
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
			DSN:  "data/notify.db?_fk=1",
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
			log.Println("Failed to read config file.")
			log.Printf("%+v\n", err)
			os.Exit(1)
		}
	}

	err = yaml.Unmarshal(data, &config)
	config.Raw = data
	if err != nil {
		log.Println("Failed to parse config file.")
		log.Printf("%+v\n", err)
		os.Exit(1)
	}
	return &config
}
