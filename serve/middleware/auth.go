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
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"body": "User ID not exist",
		})
		return
	} else {
		if user.Controller.Is(userID) {
			c.Set(UserIdKey, userID)
			c.Next()
			return
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"body": "User " + userID + " is not valid",
			})
			return
		}
	}
}
