package main

import (
	"embed"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"notify-api/db/entity"
	"notify-api/push/providers"
	"notify-api/serve/middleware"
	"notify-api/user"
	"notify-api/utils"
	"notify-api/web"
)

//go:embed "static/*"
var f embed.FS

func setup(router *gin.Engine) {
	go utils.CheckInternetConnection()

	entity.Init()
	providers.Init(router)
	user.Init()
	web.Init(&f)

	router.Use(cors.Default())
	router.Use(middleware.Auth)
}
