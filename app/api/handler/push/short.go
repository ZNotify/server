package push

import (
	"io"
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"notify-api/app/api/common"
	"notify-api/app/api/models"
	"notify-api/app/db/dao"
	"notify-api/app/manager/push"
	"notify-api/app/manager/push/enum"
	"notify-api/app/manager/push/item"
)

// Short godoc
//
//	@Summary      Send notification
//	@Id           sendMessageLite
//	@Tags         Message
//	@Description  Send notification to user_id
//	@Param        user_secret  path  string  true  "Secret of user"
//	@Param        string       body  string  true  "Message Content"
//	@Accept       plain
//	@Produce      json
//	@Success      200  {object}  common.Response[models.Message]
//	@Failure      400  {object}  common.BadRequestResponse
//	@Failure      401  {object}  common.UnauthorizedResponse
//	@Router       /{user_secret} [post]
func Short(context *common.Context) {
	data, err := io.ReadAll(context.Request.Body)
	if err != nil {
		zap.S().Errorf("read body error: %v", err)
		context.JSONError(http.StatusInternalServerError, err)
		return
	}

	if len(data) == 0 {
		zap.S().Errorf("request body is empty")
		context.JSONError(http.StatusBadRequest, errors.New("content can not be empty"))
		return
	}

	pushMsg := item.NewPushMessage(context.User, "Notification", string(data), "", enum.PriorityNormal)

	err = push.Send(context, pushMsg)
	if err != nil {
		zap.S().Errorw("send message error", "error", err)
		context.JSONError(http.StatusInternalServerError, errors.WithStack(err))
		return
	}

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
