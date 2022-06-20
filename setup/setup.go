package setup

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"notify-api/db/entity"
	"notify-api/push"
	"notify-api/push/providers"
	"notify-api/serve/handler"
	"notify-api/serve/middleware"
	"notify-api/user"
	"notify-api/web"
)

var router = gin.Default()

func New() *gin.Engine {
	checkConnection()

	entity.Init()
	providers.Init()
	user.Controller.Init()
	web.Init()

	setupRouter()
	setupMiddleware()

	return router
}

func checkConnection() {
	go func() {
		_, err := http.Get("https://www.google.com/robots.txt")
		if err != nil {
			fmt.Println("No global internet connection")
			panic(err)
		}
	}()
}

func setupMiddleware() {
	router.Use(cors.Default())
	router.Use(middleware.Auth)
}

func setupRouter() {
	router.GET("/check", handler.Check)
	router.GET("/:user_id/record", handler.Record)
	router.POST("/:user_id/send", handler.Send)
	router.DELETE("/:user_id/:id", handler.Delete)

	router.StaticFS("/fs", handler.UI)
	router.GET("/", handler.Index)

	router.GET("/alive", handler.Alive)

	err := push.Providers.RegisterRouter(router)
	if err != nil {
		panic(err)
	}
}
