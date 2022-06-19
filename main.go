package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"notify-api/serve/handler"
)

var err error

var router = gin.Default()

func main() {
	setup(router)

	router.GET("/check", handler.Check)
	router.GET("/:user_id/record", handler.Record)
	router.POST("/:user_id/send", handler.Send)
	router.DELETE("/:user_id/:id", handler.Delete)

	router.StaticFS("/fs", handler.UI)
	router.GET("/", handler.Index)

	router.GET("/alive", handler.Alive)

	err = router.Run("0.0.0.0:14444")
	if err != nil {
		panic(err)
	}
}
