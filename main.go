package main

import (
	"embed"
	"github.com/ZNotify/server/db"
	"github.com/ZNotify/server/handler"
	"github.com/ZNotify/server/push"
	"github.com/ZNotify/server/user"
	"github.com/ZNotify/server/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"io/fs"
	"net/http"
)

//go:embed static/*
var f embed.FS

var err error
var pureFS fs.FS
var router = gin.Default()

func main() {
	router.GET("/:user_id/check", handler.Check)
	router.GET("/:user_id/record", handler.Record)
	router.DELETE("/:user_id/:id", handler.Delete)
	router.POST("/:user_id/send", handler.Send)

	router.StaticFS("/fs", http.FS(pureFS))
	router.GET("/", handler.IndexWithFS(pureFS))

	router.GET("/alive", handler.Alive)

	err = router.Run("0.0.0.0:14444")
	if err != nil {
		panic("Server failed to listen.")
	}
}

func init() {
	go utils.CheckInternetConnection()
	db.Init()
	push.Init(router)
	user.Init()

	router.Use(cors.Default())

	pureFS, err = fs.Sub(f, "static")
	if err != nil {
		panic(err)
	}
}
