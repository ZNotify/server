package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notify-api/user"
)

const UserIdKey = "user_id"

func Auth(c *gin.Context) {
	userID, shouldCheck := c.Params.Get(UserIdKey)
	if !shouldCheck {
		c.Next()
		return
	} else {
		if user.IsUser(userID) {
			c.Set(UserIdKey, userID)
			c.Next()
			return
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
