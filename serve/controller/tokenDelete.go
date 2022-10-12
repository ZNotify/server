package controller

import (
	"net/http"

	"notify-api/db/model"
	"notify-api/serve/types"
)

// TokenDelete godoc
// @Summary     Delete token
// @Description Delete token of device
// @Param       user_id   path string true "user_id"
// @Param       device_id path string true "device_id"
// @Produce     json
// @Success     200 {object} types.Response[bool]
// @Router      /{user_id}/token/{device_id} [delete]
func TokenDelete(context *types.Ctx) {
	deviceId := context.Param("device_id")

	err := model.TokenUtils.Delete(context.UserID, deviceId)
	if err != nil {
		context.JSONError(http.StatusInternalServerError, err.Error())
		return
	}

	context.JSONResult(true)
}
