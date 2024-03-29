package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ZNotify/server/app/db/dao"
	"github.com/ZNotify/server/app/db/ent/generate"
)

const userSecretKey = "user_secret"
const userContextKey = "user"

func UserAuth(c *gin.Context) {
	userSecret, ok := c.Params.Get(userSecretKey)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"body": "User secret not exist",
		})
		return
	} else {
		u, ok := dao.User.GetUserBySecret(c, userSecret)
		if ok {
			c.Set(userContextKey, u)
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

func GetUser(c *gin.Context) (user *generate.User, isAuthed bool) {
	value, isAuthed := c.Get(userContextKey)
	if isAuthed {
		user = value.(*generate.User)
	}
	return
}
