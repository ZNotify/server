package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"notify-api/db/model"
)

func RecordDetail(context *gin.Context) {
	messageID := context.Param("id")

	if messageID == "" {
		context.String(http.StatusBadRequest, "Message ID can not be empty.")
		return
	}

	message, err := model.MessageUtils.Get(messageID)
	if err != nil {
		if errors.Is(err, model.ErrRecordNotFound) {
			context.String(http.StatusNotFound, "Message not found.")
			return
		}
		context.String(http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, message)
}
