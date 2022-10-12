package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"notify-api/user"
)

const UserIdKey = "user_id"

func UserAuth(c *gin.Context) {
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
