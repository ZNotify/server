package message

import (
	"net/http"

	"notify-api/app/common"
	"notify-api/app/models"
	"notify-api/db/dao"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Get godoc
//
//	@Summary      Get message record detail
//	@Id           getMessageById
//	@Tags         Message
//	@Description  Get message record detail of a message
//	@Param        user_secret  path  string  true  "Secret of user"
//	@Param        id           path  string  true  "ID of message"
//	@Produce      json
//	@Success      200  {object}  common.Response[models.Message]
//	@Failure      400  {object}  common.BadRequestResponse
//	@Failure      401  {object}  common.UnauthorizedResponse
//	@Failure      404  {object}  common.NotFoundResponse
//	@Router       /{user_secret}/message/{id} [get]
func Get(context *common.Context) {
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

	context.JSONResult(models.FromModelMessage(*message))
}
