package main

import (
	_ "net/http/pprof"

	_ "github.com/joho/godotenv/autoload"

	"github.com/ZNotify/server/cmd"
)

// Main
//
//	@title            ZNotify API
//	@version          1.0
//	@description      This is Znotify api server.
//	@path             /
//
//	@contact.name     Issues
//	@contact.url      https://github.com/ZNotify/server/issues
//
//	@tag.name         Device
//	@tag.description  Device management
//	@tag.name         User
//	@tag.description  User management
//	@tag.name         Message
//	@tag.description  Message management
//	@tag.name         Health
//	@tag.description  Health check
//	@tag.name         UI
//	@tag.description  UI for documentation and WebPush
//	@tag.name         Push
//	@tag.description  Endpoint for push service
//
//	@license.name     Apache 2.0
//	@license.url      https://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	cmd.Execute()
}
