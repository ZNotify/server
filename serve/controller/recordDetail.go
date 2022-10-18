package controller

import (
	"net/http"

	"github.com/pkg/errors"

	"notify-api/db/model"
	"notify-api/serve/types"
)

// RecordDetail godoc
// @Summary     Get message record detail
// @Description Get message record detail of a message
// @Param       user_id path string true "user_id"
// @Param       id      path string true "id"
// @Produce     json
// @Success     200 {object} types.Response[entity.Message]
// @Failure     400 {object} types.BadRequestResponse
// @Failure     401 {object} types.UnauthorizedResponse
// @Failure     404 {object} types.NotFoundResponse
// @Router      /{user_id}/{id} [get]
func RecordDetail(context *types.Ctx) {
	messageID := context.Param("id")

	if messageID == "" {
		context.JSONError(http.StatusBadRequest, errors.New("message ID can not be empty"))
		return
	}

	message, err := model.MessageUtils.Get(messageID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			context.JSONError(http.StatusNotFound, errors.New("Message not found."))
			return
		}
		context.JSONError(http.StatusInternalServerError, errors.WithStack(err))
		return
	}

	if message.UserID != context.UserID {
		context.JSONError(http.StatusUnauthorized, errors.New("You are not authorized to access this message."))
		return
	}

	context.JSONResult(message)
}
