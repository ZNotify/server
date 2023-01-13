package setup

import (
	"net/http"
	"net/http/pprof"
	"time"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"notify-api/ent/dao"
	"notify-api/server/controller/device"
	"notify-api/server/controller/message"
	"notify-api/server/controller/misc"
	"notify-api/server/controller/send"
	"notify-api/server/controller/user"
	"notify-api/setup/config"
	setupMisc "notify-api/setup/misc"
	"notify-api/setup/oauth"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"notify-api/docs"
	"notify-api/push"
	"notify-api/server/middleware"
	"notify-api/server/types"
	"notify-api/web"
)

func Setup() {
	setupMisc.RequireNetwork()
	setupMisc.RequireX64()

	dao.Init()
	push.Init()
	oauth.Init()
}

func NewRouter() *gin.Engine {
	router := gin.New()

	setupDoc(router)
	setupMiddleware(router)
	setupController(router)

	return router
}

func setupMiddleware(router *gin.Engine) {
	router.Use(cors.Default())
}

func setupDoc(router *gin.Engine) {
	if config.IsProd() {
		docs.SwaggerInfo.Schemes = append(docs.SwaggerInfo.Schemes, "https")
	} else {
		docs.SwaggerInfo.Schemes = append(docs.SwaggerInfo.Schemes, "http")
	}

	router.GET("/docs", types.WrapHandler(misc.DocRedirect))
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func setupController(router *gin.Engine) {
	router.Use(ginzap.Ginzap(zap.L(), time.RFC3339, false))
	router.Use(ginzap.RecoveryWithZap(zap.L(), true))

	router.Use(middleware.Duration)

	router.GET("/check", types.WrapHandler(user.Check))
	router.GET("/alive", types.WrapHandler(misc.Alive))

	loginGroup := router.Group("/login")
	{
		loginGroup.GET("", types.WrapHandler(user.Login))
		loginGroup.GET("/github", types.WrapHandler(user.GitHub))
	}

	userGroup := router.Group("/:user_secret")
	userGroup.Use(middleware.UserAuth)
	{
		userGroup.GET("/messages", types.WrapHandler(user.Messages))
		userGroup.GET("/message/:id", types.WrapHandler(message.Get))
		userGroup.DELETE("/message/:id", types.WrapHandler(message.Delete))

		userGroup.GET("/devices", types.WrapHandler(user.Devices))
		userGroup.PUT("/device/:device_id", types.WrapHandler(device.Create))
		userGroup.DELETE("/device/:device_id", types.WrapHandler(device.Delete))

		userGroup.POST("/send", types.WrapHandler(send.Send))
		userGroup.PUT("/send", types.WrapHandler(send.Send))

		userGroup.POST("", types.WrapHandler(send.Short))
		userGroup.PUT("", types.WrapHandler(send.Short))

		push.RegisterRouter(userGroup)
	}

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

	router.StaticFS("/fs", web.StaticHttpFS)
	router.GET("/", types.WrapHandler(misc.WebIndex))
}
