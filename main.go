package main

import (
	_ "github.com/joho/godotenv/autoload"
	"notify-api/setup"
)

var router = setup.New()

func main() {
	err := router.Run("0.0.0.0:14444")
	if err != nil {
		panic(err)
	}
}
