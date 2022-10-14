package controller

import (
	"net/http"

	"github.com/pkg/errors"

	"notify-api/db/model"
	"notify-api/serve/types"
)

// RecordDelete godoc
// @Summary     Delete message record
// @Description Delete message record with id
// @Param       user_id path string true "user_id"
// @Param       id      path string true "id"
// @Produce     json
// @Success     200 {object} types.Response[bool]
// @Failure     401 {object} types.UnauthorizedResponse
// @Router      /{user_id}/{id} [delete]
func RecordDelete(context *types.Ctx) {
	id := context.Param("id")

	err := model.MessageUtils.Delete(context.UserID, id)
	if err != nil {
		context.JSONError(http.StatusInternalServerError, errors.WithStack(err))
		return
	}

	context.JSONResult(true)
}