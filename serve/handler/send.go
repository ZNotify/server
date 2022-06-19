package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"notify-api/db"
	"notify-api/db/entity"
	"notify-api/push"
	"notify-api/serve/middleware"
	"notify-api/utils"
	"time"
)

func Send(context *gin.Context) {
	userID := context.GetString(middleware.UserIdKey)

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
		context.SecureJSON(http.StatusOK, message)
		return
	}

	err := push.Send(message)
	if err != nil {
		context.String(http.StatusInternalServerError, fmt.Sprintf("%s", err))
	}

	// Insert message record
	db.DB.Create(message)

	context.SecureJSON(http.StatusOK, message)
}
