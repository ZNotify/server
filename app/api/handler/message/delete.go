package message

import (
	"net/http"

	"github.com/ZNotify/server/app/api/common"
	"github.com/ZNotify/server/app/db/dao"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Delete godoc
//
//	@Summary      Delete message record
//	@Id           deleteMessageById
//	@Tags         Message
//	@Description  Delete message record with id
//	@Param        user_secret  path  string  true  "Secret of user"
//	@Param        id           path  string  true  "ID of message"
//	@Produce      json
//	@Success      200  {object}  common.Response[bool]
//	@Failure      401  {object}  common.UnauthorizedResponse
//	@Router       /{user_secret}/message/{id} [delete]
func Delete(context *common.Context) {
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
