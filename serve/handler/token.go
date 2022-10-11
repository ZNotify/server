package handler

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"notify-api/serve/middleware"
)

func Token(context *gin.Context) {
	userID := context.GetString(middleware.UserIdKey)

	channel := context.Param("channel")

	deviceID := context.Param("device_id")
	if len(deviceID) != 36 {
		context.String(http.StatusBadRequest, "Invalid device id")
		return
	}

	token, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}
	tokenString := string(token)

	_ = channel
	_ = tokenString
	_ = userID
	// entity.PushTokenUtils.CreateOrUpdate(userID, deviceID, tokenString)

}
