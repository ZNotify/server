package main

import (
	"embed"
	"github.com/ZNotify/server/config"
	"github.com/ZNotify/server/db"
	"github.com/ZNotify/server/handler"
	"github.com/ZNotify/server/push"
	"github.com/ZNotify/server/user"
	"github.com/ZNotify/server/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
)

//go:embed static/*
var f embed.FS

func main() {
	var err error

	config.CheckMiPushSecret()
	config.CheckFCMCredential()

	go utils.CheckInternetConnection()

	pureFs, err := fs.Sub(f, "static")
	if err != nil {
		panic(err)
	}

	db.InitDB()
	user.ReadUsers()

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/:user_id/check", handler.Check)
	router.GET("/:user_id/record", handler.Record)
	router.DELETE("/:user_id/:id", handler.Delete)
	router.POST("/:user_id/send", handler.Send)

	router.PUT("/:user_id/fcm/token", push.SetFCMToken)

	router.StaticFS("/fs", http.FS(pureFs))
	router.GET("/", func(context *gin.Context) {
		context.FileFromFS("/", http.FS(pureFs))
		// hardcode index.html, use this as a trick to get html file
		// https://github.com/golang/go/blob/a7e16abb22f1b249d2691b32a5d20206282898f2/src/net/http/fs.go#L587
	})

	err = router.Run("0.0.0.0:14444")
	if err != nil {
		panic("Server failed to listen.")
	}
}
