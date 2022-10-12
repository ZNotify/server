package controller

import (
	"net/http"

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
// @Failure     401 {string} string "Unauthorized"
// @Router      /{user_id}/{id} [delete]
func RecordDelete(context *types.Ctx) {
	id := context.Param("id")

	err := model.MessageUtils.Delete(context.UserID, id)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}

	context.JSONResult(true)
}
