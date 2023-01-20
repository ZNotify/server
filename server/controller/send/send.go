package send

import (
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"notify-api/ent/dao"
	"notify-api/push"
	"notify-api/push/enum"
	"notify-api/push/item"
	"notify-api/server/types"
	"notify-api/server/types/entity"
)

// Send godoc
//
//	@Summary		Send notification
//	@Id				sendMessage
//	@Description	Send notification to user_id
//	@Param			user_secret	path		string			true	"Secret of user"
//	@Param			title		formData	string			false	"Message Title"	default(Notification)
//	@Param			content		formData	string			true	"Message Content"
//	@Param			long		formData	string			false	"Long Message Content (optional)"
//	@Param			priority	formData	enum.Priority	false	"The priority of message"	default(normal)
//	@Produce		json
//	@Success		200	{object}	types.Response[entity.Message]
//	@Failure		400	{object}	types.BadRequestResponse
//	@Failure		401	{object}	types.UnauthorizedResponse
//	@Router			/{user_secret}/send  [post]
func Send(context *types.Ctx) {
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

	var priorityConst enum.Priority
	switch priority {
	case "low":
		priorityConst = enum.PriorityLow
	case "normal":
		priorityConst = enum.PriorityNormal
	case "high":
		priorityConst = enum.PriorityHigh
	default:
		zap.S().Infof("priority is invalid")
		context.JSONError(http.StatusBadRequest, errors.New("priority is invalid"))
		return
	}

	pushMsg := item.NewPushMessage(context.User, title, content, long, priorityConst)

	err := push.Send(context, pushMsg)
	if err != nil {
		zap.S().Errorw("send message error", "error", err)
		context.JSONError(http.StatusInternalServerError, errors.WithStack(err))
		return
	}

	// Insert message record
	msg, ok := dao.Message.CreateMessage(
		context,
		pushMsg.User,
		pushMsg.ID,
		pushMsg.Title,
		pushMsg.Content,
		pushMsg.Long,
		pushMsg.Priority,
		pushMsg.SequenceID,
		pushMsg.CreatedAt)

	if !ok {
		context.JSONError(http.StatusInternalServerError, errors.New("can not create message"))
		return
	}

	context.JSONResult(entity.FromModelMessage(*msg))
	return
}
