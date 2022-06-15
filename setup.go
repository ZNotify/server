package main

import (
	"embed"
	"github.com/ZNotify/server/db"
	"github.com/ZNotify/server/push"
	"github.com/ZNotify/server/serve/middleware"
	"github.com/ZNotify/server/user"
	"github.com/ZNotify/server/utils"
	"github.com/ZNotify/server/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//go:embed "static/*"
var f embed.FS

func setup(router *gin.Engine) {
	go utils.CheckInternetConnection()

	db.Init()
	push.Init(router)
	user.Init()
	web.Init(&f)

	router.Use(cors.Default())
	router.Use(middleware.Auth)
}
