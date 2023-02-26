package push

import (
	"io"
	"net/http"

	"notify-api/ent/dao"
	"notify-api/push"
	"notify-api/push/enum"
	"notify-api/push/item"
	"notify-api/server/types"
	"notify-api/server/types/entity"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Short godoc
//
//	@Summary		Send notification
//	@Id				sendMessageLite
//	@Tags			Message
//	@Description	Send notification to user_id
//	@Param			user_secret	path	string	true	"Secret of user"
//	@Param			string		body	string	true	"Message Content"
//	@Accept			plain
//	@Produce		json
//	@Success		200	{object}	types.Response[entity.Message]
//	@Failure		400	{object}	types.BadRequestResponse
//	@Failure		401	{object}	types.UnauthorizedResponse
//	@Router			/{user_secret} [post]
func Short(context *types.Ctx) {
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

	context.JSONResult(entity.FromModelMessage(*msg))
}
