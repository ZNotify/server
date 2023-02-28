package push

import (
	"net/http"

	"notify-api/app/common"
	"notify-api/app/manager/push"
	"notify-api/app/manager/push/enum"
	"notify-api/app/manager/push/item"
	"notify-api/app/models"
	"notify-api/db/dao"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Send godoc
//
//	@Summary      Send notification
//	@Id           sendMessage
//	@Tags         Message
//	@Description  Send notification to user_id
//	@Param        user_secret  path      string         true   "Secret of user"
//	@Param        title        formData  string         false  "Message Title"  default(Notification)
//	@Param        content      formData  string         true   "Message Content"
//	@Param        long         formData  string         false  "Long Message Content (optional)"
//	@Param        priority     formData  enum.Priority  false  "The priority of message"  default(normal)
//	@Produce      json
//	@Success      200  {object}  common.Response[models.Message]
//	@Failure      400  {object}  common.BadRequestResponse
//	@Failure      401  {object}  common.UnauthorizedResponse
//	@Router       /{user_secret}/send  [post]
func Send(context *common.Context) {
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

	context.JSONResult(models.FromModelMessage(*msg))
}
