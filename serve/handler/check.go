package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notify-api/user"
	"strconv"
)

func Check(context *gin.Context) {
	userID := context.Query("user_id")
	result := user.Controller.Is(userID)
	context.String(http.StatusOK, strconv.FormatBool(result))
	return
}
