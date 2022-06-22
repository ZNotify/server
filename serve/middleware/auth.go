package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notify-api/user"
)

const UserIdKey = "user_id"

func Auth(c *gin.Context) {
	userID, ok := c.Params.Get(UserIdKey)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	} else {
		if user.Controller.Is(userID) {
			c.Set(UserIdKey, userID)
			c.Next()
			return
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
