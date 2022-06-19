package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"notify-api/db"
	"notify-api/push"
	"notify-api/serve/middleware"
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

	message := &push.Message{
		ID:        uuid.New().String(),
		UserID:    userID,
		Title:     title,
		Content:   content,
		Long:      long,
		CreatedAt: time.Now(),
	}

	err := push.Send(message)
	if err != nil {
		context.String(http.StatusInternalServerError, fmt.Sprintf("%s", err))
	}

	// Insert message record
	db.DB.Create(message)

	context.SecureJSON(http.StatusOK, message)
}
