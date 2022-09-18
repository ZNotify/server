package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notify-api/db/model"
	"notify-api/serve/middleware"
)

func Record(context *gin.Context) {
	userID := context.GetString(middleware.UserIdKey)

	messages, err := model.MessageUtils.GetMessageInMonth(userID)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, messages)
}
