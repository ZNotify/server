package types

import (
	"log"
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

func (c *Ctx) JSONError(code int, value string) {
	switch code {
	case http.StatusBadRequest:
		c.JSON(http.StatusBadRequest, BadRequestResponse{
			Code: code,
			Body: value,
		})
	case http.StatusUnauthorized:
		c.JSON(http.StatusUnauthorized, UnauthorizedResponse{
			Code: code,
			Body: value,
		})
	case http.StatusInternalServerError:
		c.JSON(http.StatusInternalServerError, InternalServerErrorResponse{
			Code: code,
			Body: value,
		})
		log.Printf("Internal Server Error:\n%s", value)
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
