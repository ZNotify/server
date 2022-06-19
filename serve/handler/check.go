package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notify-api/user"
	"strconv"
)

func Check(context *gin.Context) {
	userID := context.Query("user_id")
	result := user.IsUser(userID)
	context.String(http.StatusOK, strconv.FormatBool(result))
	return
}
