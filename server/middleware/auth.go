package middleware

import (
	"net/http"

	"notify-api/utils/user"

	"github.com/gin-gonic/gin"
)

const UserIdKey = "user_id"
const UserKey = "user"

func UserAuth(c *gin.Context) {
	userID, ok := c.Params.Get(UserIdKey)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"body": "User ID not exist",
		})
		return
	} else {
		if user.Is(userID) {
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
