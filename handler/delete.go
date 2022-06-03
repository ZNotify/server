package handler

import (
	"errors"
	"github.com/ZNotify/server/db"
	"github.com/ZNotify/server/db/entity"
	"github.com/ZNotify/server/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Delete(context *gin.Context) {
	userID, err := utils.RequireAuth(context)
	if err != nil {
		utils.BreakOnError(context, err)
		return
	}

	id := context.Param("id")

	var message entity.Message
	result := db.DB.Where("user_id = ?", userID).
		Where("id = ?", id).
		First(&message)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		context.String(http.StatusNotFound, "Not Found")
		return
	} else {
		utils.BreakOnError(context, result.Error)
	}
	result = db.DB.Delete(&message)
	utils.BreakOnError(context, result.Error)

	context.String(http.StatusOK, "OK")
	return
}
