package push

import (
	"github.com/gin-gonic/gin"

	"github.com/ZNotify/server/app/api/common"
	"github.com/ZNotify/server/app/manager/push/interfaces"
)

func RegisterRouter(e *gin.RouterGroup) {
	for _, v := range activeSenders {
		if pv, ok := v.(interfaces.SenderWithHandler); ok {
			e.Handle(pv.HandlerMethod(), pv.HandlerPath(), common.WrapHandler(pv.Handler))
		}
	}
}
