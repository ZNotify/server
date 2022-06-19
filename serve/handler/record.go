package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notify-api/db"
	"notify-api/db/entity"
	"notify-api/serve/middleware"
	"time"
)

func Record(context *gin.Context) {
	userID := context.GetString(middleware.UserIdKey)

	var messages []entity.Message
	result := db.DB.Where("user_id = ?", userID).
		Where("created_at > ?", time.Now().AddDate(0, 0, -30)).
		Order("created_at desc").
		Find(&messages)
	if result.Error != nil {
		context.String(http.StatusInternalServerError, result.Error.Error())
		return
	}

	context.JSON(http.StatusOK, messages)
}
