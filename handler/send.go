package handler

import (
	"fmt"
	"github.com/ZNotify/server/db"
	"github.com/ZNotify/server/db/entity"
	"github.com/ZNotify/server/push"
	"github.com/ZNotify/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func Send(context *gin.Context) {
	userID, err := utils.RequireAuth(context)
	if err != nil {
		utils.BreakOnError(context, err)
		return
	}

	// string to bool
	dryRun := context.Request.URL.Query().Has("dry")
	if dryRun {
		context.String(http.StatusBadRequest, "Dry query param is not supported now.")
		return
	}

	// get notification info
	title := context.DefaultPostForm("title", "Notification")
	content := context.PostForm("content")
	long := context.PostForm("long")

	if content == "" {
		context.String(http.StatusBadRequest, "Content can not be empty.")
		return
	}

	message := &entity.Message{
		ID:        uuid.New().String(),
		UserID:    userID,
		Title:     title,
		Content:   content,
		Long:      long,
		CreatedAt: time.Now(),
	}

	if utils.IsTestInstance() {
		context.SecureJSON(http.StatusOK, message.ToGinH())
		return
	}

	err = push.Send(message)
	if err != nil {
		context.String(http.StatusInternalServerError, fmt.Sprintf("%s", err))
	}

	// Insert message record
	db.DB.Create(message)

	context.SecureJSON(http.StatusOK, message.ToGinH())
}
