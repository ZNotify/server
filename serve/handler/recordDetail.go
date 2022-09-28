package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notify-api/db/entity"
	"notify-api/db/model"
)

func RecordDetail(context *gin.Context) {
	messageID := context.Param("id")

	if messageID == "" {
		context.String(http.StatusBadRequest, "Message ID can not be empty.")
		return
	}

	message, err := model.MessageUtils.GetOrEmpty(messageID)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	if message == entity.EmptyMessage {
		context.String(http.StatusNotFound, "Message not found.")
		return
	}

	context.JSON(http.StatusOK, message)
}
