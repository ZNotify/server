package handler

import (
	"github.com/ZNotify/server/db"
	"github.com/ZNotify/server/db/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Record(context *gin.Context) {
	userID := context.GetString("user_id")

	var messages []entity.Message
	result := db.DB.Where("user_id = ?", userID).
		Where("created_at > ?", time.Now().AddDate(0, 0, -30)).
		Order("created_at desc").
		Find(&messages)
	if result.Error != nil {
		context.String(http.StatusInternalServerError, result.Error.Error())
		return
	}

	var ret []gin.H
	for i := range messages {
		ret = append(ret, messages[i].ToGinH())
	}
	context.JSON(http.StatusOK, ret)
}
