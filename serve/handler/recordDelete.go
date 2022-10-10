package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notify-api/db/model"
	"notify-api/serve/middleware"
)

func RecordDelete(context *gin.Context) {
	userID := context.GetString(middleware.UserIdKey)

	id := context.Param("id")

	err := model.MessageUtils.Delete(userID, id)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}

	context.String(http.StatusOK, "OK")
}