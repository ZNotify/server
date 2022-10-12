package types

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"notify-api/serve/middleware"
)

type Ctx struct {
	*gin.Context
	UserID   string
	IsAuthed bool
}

func (c *Ctx) JSONResult(value any) {
	ret := Response[any]{
		Code: http.StatusOK,
		Body: value,
	}
	c.JSON(http.StatusOK, ret)
}

func WrapHandler(handler func(*Ctx)) gin.HandlerFunc {
	return func(context *gin.Context) {
		userID := ""
		value, isAuthed := context.Get(middleware.UserIdKey)
		if isAuthed {
			userID, _ = value.(string)
		}

		ctx := new(Ctx)
		ctx.Context = context
		ctx.UserID = userID
		ctx.IsAuthed = isAuthed

		handler(ctx)
	}
}
