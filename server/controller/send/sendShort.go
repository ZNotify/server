package send

import (
	"io"
	"net/http"
	"time"

	"notify-api/ent/dao"
	"notify-api/push"
	"notify-api/push/item"
	"notify-api/server/types"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Short godoc
//
//	@Summary		Send notification
//	@Description	Send notification to user_id
//	@Param			user_secret	path	string	true	"Secret of user"
//	@Param			string		body	string	true	"Message Content"
//	@Accept			plain
//	@Produce		json
//	@Success		200	{object}	types.Response[types.Message]
//	@Failure		400	{object}	types.BadRequestResponse
//	@Failure		401	{object}	types.UnauthorizedResponse
//	@Router			/{user_secret} [post]
//	@Router			/{user_secret} [put]
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

	pushMsg := &item.PushMessage{
		ID:        uuid.New(),
		User:      context.User,
		Title:     "Notification",
		Content:   string(data),
		Long:      "",
		Priority:  item.PriorityNormal,
		CreatedAt: time.Now(),
	}

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
		pushMsg.CreatedAt)

	if !ok {
		context.JSONError(http.StatusInternalServerError, errors.New("can not create message"))
		return
	}

	context.JSONResult(types.FromModelMessage(*msg))
	return
}
