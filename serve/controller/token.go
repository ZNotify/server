package controller

import (
	"net/http"

	"notify-api/db/model"
	"notify-api/push"
	"notify-api/serve/types"
	"notify-api/utils"
)

// Token godoc
// @Summary     Create or update token
// @Description Create or update token of device
// @Param       user_id   path     string true "user_id"
// @Param       device_id path     string true "device_id"
// @Param       channel   formData string true "channelm should be a valid UUID"
// @Param       token     formData string true "token"
// @Produce     json
// @Success     200 {object} types.Response[bool]
// @Failure     400 {object} types.BadRequestResponse
// @Failure     401 {object} types.UnauthorizedResponse
// @Router      /{user_id}/token/{device_id} [put]
func Token(context *types.Ctx) {
	deviceID := context.Param("device_id")
	if !utils.IsUUID(deviceID) {
		context.JSONError(http.StatusBadRequest, "Invalid device id")
		return
	}

	channel := context.PostForm("channel")
	if !push.Senders.Has(channel) {
		context.JSONError(http.StatusBadRequest, "Invalid channel")
		return
	}

	token := context.PostForm("token")

	_, err := model.TokenUtils.CreateOrUpdate(context.UserID, deviceID, channel, token)
	if err != nil {
		context.JSONError(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSONResult(true)
}
