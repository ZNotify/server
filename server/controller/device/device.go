package device

import (
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"notify-api/ent/dao"
	"notify-api/push"
	"notify-api/server/types"
	"notify-api/utils"
)

// Device godoc
//
//	@Summary		Create or update device
//	@Description	Create or update device information
//	@Param			user_secret	path		string	true	"Secret of user"
//	@Param			device_id	path		string	true	"device_id should be a valid UUID"
//	@Param			channel		formData	string	true	"channel can be used."	Enums(TelegramHost, WebSocketHost, FCM, WebPush, WNS)
//	@Param			device_name	formData	string	false	"device name"
//	@Param			device_meta	formData	string	false	"additional device meta"
//	@Param			token		formData	string	false	"channel token"
//	@Produce		json
//	@Success		200	{object}	types.Response[bool]
//	@Failure		400	{object}	types.BadRequestResponse
//	@Failure		401	{object}	types.UnauthorizedResponse
//	@Router			/{user_id}/token/{device_id} [put]
func Device(context *types.Ctx) {
	deviceID := context.Param("device_id")
	if !utils.IsUUID(deviceID) {
		zap.S().Infof("device id %s is not a valid UUID", deviceID)
		context.JSONError(http.StatusBadRequest, errors.New("client_id should be a valid UUID"))
		return
	}

	channel := context.PostForm("channel")
	if !push.IsValid(channel) {
		zap.S().Infof("channel %s is not supported", channel)
		context.JSONError(http.StatusBadRequest, errors.New("channel is not supported"))
		return
	}

	token := context.PostForm("token")
	deviceName := context.PostForm("device_name")
	deviceMeta := context.PostForm("device_meta")

	_, ok := dao.Device.EnsureDevice(context,
		deviceID,
		channel,
		token,
		deviceName,
		deviceMeta,
		context.User,
	)

	context.JSONResult(ok)
}
