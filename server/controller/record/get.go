package record

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"notify-api/ent/dao"
	"notify-api/server/types"
)

// Get godoc
//
//		@Summary		Get message record detail
//	 @Id				record.get
//		@Description	Get message record detail of a message
//		@Param			user_secret	path	string	true	"Secret of user"
//		@Param			id			path	string	true	"ID of message"
//		@Produce		json
//		@Success		200	{object}	types.Response[types.Message]
//		@Failure		400	{object}	types.BadRequestResponse
//		@Failure		401	{object}	types.UnauthorizedResponse
//		@Failure		404	{object}	types.NotFoundResponse
//		@Router			/{user_secret}/{id} [get]
func Get(context *types.Ctx) {
	messageID, err := uuid.Parse(context.Param("id"))

	if err != nil {
		zap.S().Infof("can not parse message id %s to uuid", context.Param("id"))
		context.JSONError(http.StatusBadRequest, errors.Wrap(err, "can not parse message id"))
		return
	}

	message, ok := dao.Message.GetUserMessage(context, context.User, messageID)
	if !ok {
		context.JSONError(http.StatusInternalServerError, errors.New("can not get message"))
	}

	if message == nil {
		context.JSONError(http.StatusNotFound, errors.New("message not found"))
		return
	}

	context.JSONResult(types.FromModelMessage(*message))
}
