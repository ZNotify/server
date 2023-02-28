package push

import (
	"github.com/gin-gonic/gin"

	"notify-api/app/common"
	"notify-api/app/manager/push/interfaces"
)

func RegisterRouter(e *gin.RouterGroup) {
	for _, v := range activeSenders {
		if pv, ok := v.(interfaces.SenderWithHandler); ok {
			e.Handle(pv.HandlerMethod(), pv.HandlerPath(), common.WrapHandler(pv.Handler))
		}
	}
}
