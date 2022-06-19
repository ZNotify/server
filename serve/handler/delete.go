package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"notify-api/db"
	"notify-api/db/entity"
	"notify-api/serve/middleware"
)

func Delete(context *gin.Context) {
	userID := context.GetString(middleware.UserIdKey)

	id := context.Param("id")

	var message entity.Message
	result := db.DB.Where("user_id = ?", userID).
		Where("id = ?", id).
		First(&message)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			context.String(http.StatusNotFound, "Not Found")
			return
		} else {
			context.String(http.StatusInternalServerError, result.Error.Error())
			return
		}
	}

	result = db.DB.Delete(&message)
	if result.Error != nil {
		context.String(http.StatusInternalServerError, result.Error.Error())
		return
	}

	context.String(http.StatusOK, "OK")
	return
}
