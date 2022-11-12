package controller

import (
	"io"
	"net/http"
	"time"

	"notify-api/push"
	pushTypes "notify-api/push/types"
	"notify-api/serve/types"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// SendShort godoc
//
//	@Summary      Send notification
//	@Description  Send notification to user_id
//	@Param        user_id  path  string  true  "user_id"
//	@Param        string   body  string  true  "content"
//	@Accept       plain
//	@Produce      json
//	@Success      200  {object}  types.Response[entity.Message]
//	@Failure      400  {object}  types.BadRequestResponse
//	@Failure      401  {object}  types.UnauthorizedResponse
//	@Router       /{user_id} [post]
//	@Router       /{user_id} [put]
func SendShort(context *types.Ctx) {
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

	pushMsg := &pushTypes.Message{
		ID:        uuid.New().String(),
		UserID:    context.UserID,
		Title:     "Notification",
		Content:   string(data),
		Long:      "",
		CreatedAt: time.Now(),
	}

	err = push.Send(pushMsg)
	if err != nil {
		zap.S().Errorw("send message error", "error", err)
		context.JSONError(http.StatusInternalServerError, errors.WithStack(err))
		return
	}

	context.JSONResult(pushMsg)
	return
}
