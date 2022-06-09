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
	_ "github.com/joho/godotenv/autoload"
	"io/fs"
	"net/http"
)

//go:embed static/*
var f embed.FS

func main() {
	var err error

	go utils.CheckInternetConnection()

	config.CheckMiPushSecret()
	config.CheckFCMCredential()
	config.CheckWebPushCert()

	push.InitMiPushClient()
	push.InitFCMClient()
	push.InitWebPushOption()

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
	router.PUT("/:user_id/web/sub", push.SetWebPushSubscription)

	router.StaticFS("/fs", http.FS(pureFs))
	router.GET("/alive", handler.Alive)
	router.GET("/", handler.IndexWithFS(pureFs))

	err = router.Run("0.0.0.0:14444")
	if err != nil {
		panic("Server failed to listen.")
	}
}
