package controller

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"notify-api/db/util"
	"notify-api/push"
	pushEntity "notify-api/push/entity"
	serveTypes "notify-api/server/types"
)

// Send godoc
//
//	@Summary		Send notification
//	@Description	Send notification to user_id
//	@Param			user_id		path		string				true	"user_id"
//	@Param			title		formData	string				false	"title"	default(Notification)
//	@Param			content		formData	string				true	"content"
//	@Param			long		formData	string				false	"long"
//	@Param			priority	formData	pushEntity.Priority	false	"priority"	default("normal")
//	@Produce		json
//	@Success		200	{object}	serveTypes.Response[serveTypes.Message]
//	@Failure		400	{object}	serveTypes.BadRequestResponse
//	@Failure		401	{object}	serveTypes.UnauthorizedResponse
//	@Router			/{user_id}/send  [post]
//	@Router			/{user_id}/send  [put]
func Send(context *serveTypes.Ctx) {
	// get notification info
	title := context.DefaultPostForm("title", "Notification")
	content := context.PostForm("content")
	long := context.PostForm("long")
	priority := context.DefaultPostForm("priority", "normal")

	if content == "" {
		zap.S().Infof("content is empty")
		context.JSONError(http.StatusBadRequest, errors.New("content can not be empty"))
		return
	}

	var priorityConst pushEntity.Priority
	switch priority {
	case "low":
		priorityConst = pushEntity.PriorityLow
	case "normal":
		priorityConst = pushEntity.PriorityNormal
	case "high":
		priorityConst = pushEntity.PriorityHigh
	default:
		zap.S().Infof("priority is invalid")
		context.JSONError(http.StatusBadRequest, errors.New("priority is invalid"))
		return
	}

	pushMsg := &pushEntity.PushMessage{
		MessageID: uuid.New().String(),
		UserID:    context.UserID,
		Title:     title,
		Content:   content,
		Long:      long,
		Priority:  priorityConst,
		CreatedAt: time.Now(),
	}

	err := push.Send(pushMsg)
	if err != nil {
		zap.S().Errorw("send message error", "error", err)
		context.JSONError(http.StatusInternalServerError, errors.WithStack(err))
		return
	}

	// Insert message record
	msg, err := util.MessageUtil.Add(
		pushMsg.MessageID,
		pushMsg.UserID,
		pushMsg.Title,
		pushMsg.Content,
		pushMsg.Long,
		pushMsg.Priority,
	)

	if err != nil {
		zap.S().Errorw("save message error", "error", err)
		context.JSONError(http.StatusInternalServerError, errors.WithStack(err))
		return
	}

	context.JSONResult(serveTypes.FromModelMessage(msg))
	return
}
