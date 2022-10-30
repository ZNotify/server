package types

import (
	"fmt"
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

func (c *Ctx) JSONError(code int, err error) {
	switch code {
	case http.StatusBadRequest:
		c.JSON(http.StatusBadRequest, BadRequestResponse{
			Code: code,
			Body: err.Error(),
		})
	case http.StatusUnauthorized:
		c.JSON(http.StatusUnauthorized, UnauthorizedResponse{
			Code: code,
			Body: err.Error(),
		})
	case http.StatusInternalServerError:
		errString := fmt.Sprintf("%+v", err)
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse{
			Code: code,
			Body: errString,
		})
	}
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
