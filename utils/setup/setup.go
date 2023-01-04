package setup

import (
	"net/http"
	"net/http/pprof"
	"time"

	"notify-api/ent/db"
	"notify-api/server/controller/misc"
	"notify-api/server/controller/record"
	"notify-api/server/controller/send"
	"notify-api/server/controller/user"
	"notify-api/utils/config"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

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

	router.GET("/check", types.WrapHandler(user.Check))
	router.GET("/alive", types.WrapHandler(misc.Alive))

	userGroup := router.Group("/:user_secret")
	userGroup.Use(middleware.UserAuth)
	{
		userGroup.GET("/record", types.WrapHandler(record.Record))
		userGroup.GET("/:id", types.WrapHandler(record.Detail))
		userGroup.DELETE("/:id", types.WrapHandler(record.Delete))

		userGroup.POST("/send", types.WrapHandler(send.Send))
		userGroup.PUT("/send", types.WrapHandler(send.Send))

		userGroup.POST("", types.WrapHandler(send.SendShort))
		userGroup.PUT("", types.WrapHandler(send.SendShort))

		userGroup.PUT("/device/:device_id", types.WrapHandler(controller.Token))
		userGroup.DELETE("/device/:device_id", types.WrapHandler(controller.TokenDelete))

		push.RegisterRouter(userGroup)
	}

	debugGroup := router.Group("/debug")
	debugGroup.GET("/pprof/*pprof", gin.WrapH(http.HandlerFunc(pprof.Index)))

	router.GET("/docs", types.WrapHandler(misc.DocRedirect))
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.StaticFS("/fs", web.StaticHttpFS)
	router.GET("/", types.WrapHandler(misc.WebIndex))
}
