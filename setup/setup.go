package setup

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"notify-api/db"
	"notify-api/docs"
	"notify-api/push"
	"notify-api/serve/controller"
	"notify-api/serve/middleware"
	"notify-api/serve/types"
	"notify-api/user"
	"notify-api/web"
)

var router = gin.Default()

func New() *gin.Engine {
	checkConnection()

	db.Init()
	user.Controller.Init()
	push.Senders.Init()

	setupDoc()

	setupMiddleware()
	setupRouter()

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
}

func setupDoc() {
	if gin.Mode() == gin.ReleaseMode {
		docs.SwaggerInfo.Schemes = append(docs.SwaggerInfo.Schemes, "https")
	} else {
		docs.SwaggerInfo.Schemes = append(docs.SwaggerInfo.Schemes, "http")
	}

}

func setupRouter() {
	router.GET("/check", types.WrapHandler(controller.Check))
	router.GET("/alive", types.WrapHandler(controller.Alive))

	userGroup := router.Group("/:user_id")
	userGroup.Use(middleware.UserAuth)
	{
		userGroup.GET("/record", types.WrapHandler(controller.Record))
		userGroup.GET("/:id", types.WrapHandler(controller.RecordDetail))
		userGroup.DELETE("/:id", types.WrapHandler(controller.RecordDelete))

		userGroup.POST("/send", types.WrapHandler(controller.Send))
		userGroup.PUT("/send", types.WrapHandler(controller.Send))

		userGroup.PUT("/token/:device_id", types.WrapHandler(controller.Token))
		userGroup.DELETE("/token/:device_id", types.WrapHandler(controller.TokenDelete))
	}

	router.GET("/docs", types.WrapHandler(controller.DocIndex))
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.StaticFS("/fs", web.StaticHttpFS)
	router.GET("/", types.WrapHandler(controller.WebIndex))

	err := push.Senders.RegisterRouter(router)
	if err != nil {
		panic(err)
	}
}
