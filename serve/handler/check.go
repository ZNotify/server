package handler

import (
	"github.com/ZNotify/server/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Check(context *gin.Context) {
	userID := context.Query("user_id")
	result := user.IsUser(userID)
	context.String(http.StatusOK, strconv.FormatBool(result))
	return
}
