package setup

import (
	"net/http"
	"time"

	"notify-api/utils/config"
	"notify-api/utils/user"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	"notify-api/db"
	"notify-api/docs"
	"notify-api/push"
	"notify-api/server/controller"
	"notify-api/server/middleware"
	"notify-api/server/types"
	"notify-api/web"
)

func New() *gin.Engine {
	router := gin.New()

	checkConnection()

	db.Init()
	user.Init()
	push.Init()

	setupDoc()
	setupMiddleware(router)
	setupRouter(router)

	return router
}

func checkConnection() {
	if !config.IsProd() {
		zap.S().Info("Skip connection check in non-production mode")
		return
	}

	go func() {
		_, err := http.Get("https://www.google.com/robots.txt")
		if err != nil {
			zap.S().Panicf("Failed to connect to internet: %v", err)
		}
	}()
}

func setupMiddleware(router *gin.Engine) {
	router.Use(cors.Default())
}

func setupDoc() {
	if config.IsProd() {
		docs.SwaggerInfo.Schemes = append(docs.SwaggerInfo.Schemes, "https")
	} else {
		docs.SwaggerInfo.Schemes = append(docs.SwaggerInfo.Schemes, "http")
	}
}

func setupRouter(router *gin.Engine) {
	router.Use(ginzap.Ginzap(zap.L(), time.RFC3339, false))
	router.Use(ginzap.RecoveryWithZap(zap.L(), true))

	router.Use(middleware.Duration)

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

		userGroup.POST("", types.WrapHandler(controller.SendShort))
		userGroup.PUT("", types.WrapHandler(controller.SendShort))

		userGroup.PUT("/token/:device_id", types.WrapHandler(controller.Token))
		userGroup.DELETE("/token/:device_id", types.WrapHandler(controller.TokenDelete))

		push.RegisterRouter(userGroup)
	}

	router.GET("/docs", types.WrapHandler(controller.DocRedirect))
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.StaticFS("/fs", web.StaticHttpFS)
	router.GET("/", types.WrapHandler(controller.WebIndex))
}
