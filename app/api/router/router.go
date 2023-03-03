package router

import (
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	"github.com/ZNotify/server/app/api/common"
	"github.com/ZNotify/server/app/api/handler/device"
	"github.com/ZNotify/server/app/api/handler/message"
	"github.com/ZNotify/server/app/api/handler/misc"
	"github.com/ZNotify/server/app/api/handler/push"
	"github.com/ZNotify/server/app/api/handler/user"
	"github.com/ZNotify/server/app/api/middleware"
	"github.com/ZNotify/server/app/api/web"
	"github.com/ZNotify/server/app/global"
	pushManager "github.com/ZNotify/server/app/manager/push"
	"github.com/ZNotify/server/docs"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	setupDoc(router)
	setupApi(router)
	setupDebug(router)
	setupStatic(router)

	return router
}

func setupDoc(router *gin.Engine) {
	if global.IsProd() {
		docs.SwaggerInfo.Schemes = append(docs.SwaggerInfo.Schemes, "https")
	} else {
		docs.SwaggerInfo.Schemes = append(docs.SwaggerInfo.Schemes, "http")
	}

	router.GET("/docs", common.WrapHandler(misc.DocRedirect))
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func setupApi(router *gin.Engine) {
	router.Use(ginzap.Ginzap(zap.L(), time.RFC3339, false))
	router.Use(ginzap.RecoveryWithZap(zap.L(), true))
	router.Use(cors.Default())

	router.Use(middleware.ServerTiming)

	router.GET("/check", common.WrapHandler(user.Check))
	router.GET("/alive", common.WrapHandler(misc.Alive))
	router.GET("/webpush", common.WrapHandler(push.WebPush))

	loginGroup := router.Group("/login")
	{
		loginGroup.GET("", common.WrapHandler(user.Login))
		loginGroup.GET("/github", common.WrapHandler(user.GitHub))
	}

	userGroup := router.Group("/:user_secret")
	userGroup.Use(middleware.UserAuth)
	{
		userGroup.GET("/messages", common.WrapHandler(user.Messages))
		userGroup.GET("/message/:id", common.WrapHandler(message.Get))
		userGroup.DELETE("/message/:id", common.WrapHandler(message.Delete))

		userGroup.GET("/devices", common.WrapHandler(user.Devices))
		userGroup.PUT("/device/:device_id", common.WrapHandler(device.Create))
		userGroup.DELETE("/device/:device_id", common.WrapHandler(device.Delete))

		userGroup.POST("/send", common.WrapHandler(push.Send))
		userGroup.POST("", common.WrapHandler(push.Short))

		pushManager.RegisterRouter(userGroup)
	}
}

func setupDebug(router *gin.Engine) {
	debugGroup := router.Group("/debug/pprof")
	{
		debugGroup.GET("/", gin.WrapH(http.HandlerFunc(pprof.Index)))
		debugGroup.GET("/cmdline", gin.WrapH(http.HandlerFunc(pprof.Cmdline)))
		debugGroup.GET("/profile", gin.WrapH(http.HandlerFunc(pprof.Profile)))
		debugGroup.POST("/symbol", gin.WrapH(http.HandlerFunc(pprof.Symbol)))
		debugGroup.GET("/symbol", gin.WrapH(http.HandlerFunc(pprof.Symbol)))
		debugGroup.GET("/trace", gin.WrapH(http.HandlerFunc(pprof.Trace)))
		debugGroup.GET("/allocs", gin.WrapH(http.HandlerFunc(pprof.Handler("allocs").ServeHTTP)))
		debugGroup.GET("/block", gin.WrapH(http.HandlerFunc(pprof.Handler("block").ServeHTTP)))
		debugGroup.GET("/goroutine", gin.WrapH(http.HandlerFunc(pprof.Handler("goroutine").ServeHTTP)))
		debugGroup.GET("/heap", gin.WrapH(http.HandlerFunc(pprof.Handler("heap").ServeHTTP)))
		debugGroup.GET("/mutex", gin.WrapH(http.HandlerFunc(pprof.Handler("mutex").ServeHTTP)))
		debugGroup.GET("/threadcreate", gin.WrapH(http.HandlerFunc(pprof.Handler("threadcreate").ServeHTTP)))
	}
}

func setupStatic(router *gin.Engine) {
	router.StaticFS("/web", web.StaticHttpFS)
	router.GET("/", common.WrapHandler(misc.WebIndex))
}
