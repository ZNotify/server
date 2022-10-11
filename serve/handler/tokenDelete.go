package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notify-api/db/model"
	"notify-api/serve/middleware"
)

func TokenDelete(context *gin.Context) {
	userId := context.GetString(middleware.UserIdKey)

	deviceId := context.Param("device_id")

	err := model.TokenUtils.Delete(userId, deviceId)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}

	context.String(http.StatusOK, "OK")
}
