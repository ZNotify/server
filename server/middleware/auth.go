package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"notify-api/ent/dao"
)

const UserSecretKey = "user_secret"
const UserKey = "user"

func UserAuth(c *gin.Context) {
	userSecret, ok := c.Params.Get(UserSecretKey)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"body": "User secret not exist",
		})
		return
	} else {
		u, ok := dao.User.GetUserBySecret(c, userSecret)
		if ok {
			c.Set(UserKey, u)
			c.Next()
			return
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"body": "User secret " + userSecret + " is not valid",
			})
			return
		}
	}
}
