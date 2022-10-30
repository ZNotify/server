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
	"errors"
	_ "github.com/joho/godotenv/autoload"
	"notify-api/log"
	_ "notify-api/log"
)

//var router = setup.New()

func main() {
	log.Errorf("%+v", errors.New("test"))

	//err := router.Run("0.0.0.0:14444")
	//if err != nil {
	//	panic(err)
	//}
}
