package message

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"notify-api/ent/dao"
	"notify-api/server/types"
)

// Delete godoc
//
//	@Summary		Delete message record
//	@Id				record.delete
//	@Description	Delete message record with id
//	@Param			user_secret	path	string	true	"Secret of user"
//	@Param			id			path	string	true	"ID of message"
//	@Produce		json
//	@Success		200	{object}	types.Response[bool]
//	@Failure		401	{object}	types.UnauthorizedResponse
//	@Router			/{user_secret}/message/{id} [delete]
func Delete(context *types.Ctx) {
	id, err := uuid.Parse(context.Param("id"))
	if err != nil {
		zap.S().Infof("can not parse message id %s to uuid", context.Param("id"))
		context.JSONError(http.StatusBadRequest, errors.Wrap(err, "can not parse message id"))
		return
	}

	row, ok := dao.Message.DeleteMessageByID(context, context.User, id)
	if !ok {
		context.JSONError(http.StatusInternalServerError, errors.New("can not delete message"))
	}

	if row == 0 {
		context.JSONError(http.StatusNotFound, errors.New("message not found"))
		return
	} else {
		context.JSONResult(true)
	}
}
