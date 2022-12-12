package push

import (
	"github.com/gin-gonic/gin"

	pushTypes "notify-api/push/types"
	serveTypes "notify-api/server/types"
)

func RegisterRouter(e *gin.RouterGroup) {
	for _, v := range activeSenders {
		if pv, ok := v.(pushTypes.SenderWithHandler); ok {
			e.Handle(pv.HandlerMethod(), pv.HandlerPath(), serveTypes.WrapHandler(pv.Handler))
		}
	}
}
