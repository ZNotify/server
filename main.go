package main

import (
	_ "net/http/pprof"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"

	"notify-api/cmd"
	"notify-api/server/setup/log"
)

// @title			Notify API
// @version		1.0
// @description	This is Znotify api server.
// @path			/
//
// @contact.name	Issues
// @contact.url	https://github.com/ZNotify/server/issues
//
// @license.name	Apache 2.0
// @license.url	https://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	log.Init()
	err := cmd.App.Run(os.Args)
	if err != nil {
		zap.S().Fatalf("Failed to run app: %+v", err)
	}
}
