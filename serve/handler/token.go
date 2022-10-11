package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notify-api/db/model"
	"notify-api/push"
	"notify-api/serve/middleware"
	"notify-api/utils"
)

func Token(context *gin.Context) {
	userID := context.GetString(middleware.UserIdKey)

	deviceID := context.Param("device_id")
	if !utils.IsUUID(deviceID) {
		context.String(http.StatusBadRequest, "Invalid device id")
		return
	}

	channel := context.PostForm("channel")
	if !push.Senders.Has(channel) {
		context.String(http.StatusBadRequest, "Invalid channel")
		return
	}

	token := context.PostForm("token")

	_, err := model.TokenUtils.CreateOrUpdate(userID, deviceID, channel, token)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	context.String(http.StatusOK, "OK")
}
