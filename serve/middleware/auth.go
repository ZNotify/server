package middleware

import (
	"github.com/ZNotify/server/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(c *gin.Context) {
	userID, shouldCheck := c.Params.Get("user_id")
	if !shouldCheck {
		c.Next()
		return
	} else {
		if user.IsUser(userID) {
			c.Next()
			return
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
