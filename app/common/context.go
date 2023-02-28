package common

import (
	"fmt"
	"net/http"

	"notify-api/app/middleware/serverTiming"
	"notify-api/db/ent/generate"

	"github.com/gin-gonic/gin"

	"notify-api/app/middleware"
)

type Context struct {
	*gin.Context
	User     *generate.User
	IsAuthed bool
	Timing   *serverTiming.Timing
}

func (c *Context) JSONResult(value any) {
	ret := Response[any]{
		Code: http.StatusOK,
		Body: value,
	}
	c.JSON(http.StatusOK, ret)
}

func (c *Context) JSONError(code int, err error) {
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

func WrapHandler(handler func(*Context)) gin.HandlerFunc {
	return func(context *gin.Context) {
		user, isAuthed := middleware.GetUser(context)
		timing := middleware.GetTiming(context)

		ctx := &Context{
			Context:  context,
			User:     user,
			IsAuthed: isAuthed,
			Timing:   timing,
		}

		handler(ctx)
	}
}
