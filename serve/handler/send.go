package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"notify-api/db/model"
	. "notify-api/push"
	"notify-api/push/types"
	"notify-api/serve/middleware"
	"sync"
	"time"
)

var lock = sync.RWMutex{}

func Send(context *gin.Context) {
	userID := context.GetString(middleware.UserIdKey)

	// get notification info
	title := context.DefaultPostForm("title", "Notification")
	content := context.PostForm("content")
	long := context.PostForm("long")

	if content == "" {
		context.String(http.StatusBadRequest, "Content can not be empty.")
		return
	}

	pushMsg := &types.Message{
		ID:        uuid.New().String(),
		UserID:    userID,
		Title:     title,
		Content:   content,
		Long:      long,
		CreatedAt: time.Now(),
	}

	err := Senders.Send(pushMsg)
	if err != nil {
		context.String(http.StatusInternalServerError, fmt.Sprintf("%s", err))
	}

	// Insert message record
	lock.Lock()
	time.Sleep(1 * time.Nanosecond)
	// a trick to generate different timestamp for different message
	// FIXME: use a increasing counter to generate different id
	msg, err := model.MessageUtils.Add(
		pushMsg.ID,
		pushMsg.UserID,
		pushMsg.Title,
		pushMsg.Content,
		pushMsg.Long,
		time.Now())
	lock.Unlock()

	if err != nil {
		context.String(http.StatusInternalServerError, fmt.Sprintf("%s", err))
	}

	context.JSON(http.StatusOK, msg)
}
