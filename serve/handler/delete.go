package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notify-api/db/entity"
	"notify-api/serve/middleware"
)

func Delete(context *gin.Context) {
	userID := context.GetString(middleware.UserIdKey)

	id := context.Param("id")

	err := entity.MessageUtils.Delete(userID, id)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}

	context.String(http.StatusOK, "OK")
}
