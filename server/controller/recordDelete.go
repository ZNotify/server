package controller

import (
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"notify-api/ent/dao"
	"notify-api/server/types"
)

// RecordDelete godoc
//
//	@Summary		Delete message record
//	@Description	Delete message record with id
//	@Param			user_id	path	string	true	"user_id"
//	@Param			id		path	string	true	"id"
//	@Produce		json
//	@Success		200	{object}	types.Response[bool]
//	@Failure		401	{object}	types.UnauthorizedResponse
//	@Router			/{user_id}/{id} [delete]
func RecordDelete(context *types.Ctx) {
	id := context.Param("id")

	err := dao.Message.Delete(context.UserID, id)
	if err != nil {
		zap.S().Errorw("delete message error", "error", err)
		context.JSONError(http.StatusInternalServerError, errors.WithStack(err))
		return
	}

	context.JSONResult(true)
}
