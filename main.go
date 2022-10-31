// @title       Notify API
// @version     1.0
// @description This is Znotify api server.
// @path        /

// @contact.name Issues
// @contact.url  https://github.com/ZNotify/server/issues

// @license.name Apache 2.0
// @license.url  https://www.apache.org/licenses/LICENSE-2.0.html
package main

import (
	_ "github.com/joho/godotenv/autoload"
	"log"
	"notify-api/cmd"
	"os"
)

func main() {
	err := cmd.App.Run(os.Args)
	if err != nil {
		log.Fatalf("Failed to run app: %v", err)
	}
}
