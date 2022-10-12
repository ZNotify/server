package controller

import (
	"errors"
	"net/http"

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
		context.JSONError(http.StatusBadRequest, "Message ID can not be empty.")
		return
	}

	message, err := model.MessageUtils.Get(messageID)
	if err != nil {
		if errors.Is(err, model.ErrRecordNotFound) {
			context.JSONError(http.StatusNotFound, "Message not found.")
			return
		}
		context.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	if message.UserID != context.UserID {
		context.JSONError(http.StatusUnauthorized, "You are not the owner of this message.")
		return
	}

	context.JSONResult(message)
}
