package controller

import (
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"notify-api/db/util"
	"notify-api/server/types"
)

// RecordDetail godoc
//
//	@Summary		Get message record detail
//	@Description	Get message record detail of a message
//	@Param			user_id	path	string	true	"user_id"
//	@Param			id		path	string	true	"id"
//	@Produce		json
//	@Success		200	{object}	types.Response[types.Message]
//	@Failure		400	{object}	types.BadRequestResponse
//	@Failure		401	{object}	types.UnauthorizedResponse
//	@Failure		404	{object}	types.NotFoundResponse
//	@Router			/{user_id}/{id} [get]
func RecordDetail(context *types.Ctx) {
	messageID := context.Param("id")

	if messageID == "" {
		zap.S().Infof("message id is empty")
		context.JSONError(http.StatusBadRequest, errors.New("message ID can not be empty"))
		return
	}

	message, err := util.MessageUtil.Get(messageID)
	if err != nil {
		if errors.Is(err, util.ErrNotFound) {
			zap.S().Infof("message %s not found", messageID)
			context.JSONError(http.StatusNotFound, errors.New("Message not found."))
			return
		}
		zap.S().Errorw("get message error", "error", err)
		context.JSONError(http.StatusInternalServerError, errors.WithStack(err))
		return
	}

	if message.UserID != context.UserID {
		zap.S().Infof("message %s not belong to user %s", messageID, context.UserID)
		context.JSONError(http.StatusUnauthorized, errors.New("You are not authorized to access this message."))
		return
	}

	context.JSONResult(types.FromModelMessage(message))
}
