package device

import (
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"notify-api/ent/dao"
	"notify-api/push"
	"notify-api/push/enum"
	pushTypes "notify-api/push/types"
	"notify-api/server/types"
	"notify-api/utils"
)

// Create godoc
//
//	@Summary		Create or update device
//	@Id				device.create
//	@Description	Create or update device information
//	@Param			user_secret	path		string		true	"Secret of user"
//	@Param			device_id	path		string		true	"device_id should be a valid UUID"
//	@Param			channel		formData	enum.Sender	true	"channel can be used."
//	@Param			device_name	formData	string		false	"device name"
//	@Param			device_meta	formData	string		false	"additional device meta"
//	@Param			token		formData	string		false	"channel token"
//	@Produce		json
//	@Success		200	{object}	types.Response[bool]
//	@Failure		400	{object}	types.BadRequestResponse
//	@Failure		401	{object}	types.UnauthorizedResponse
//	@Router			/{user_secret}/device/{device_id} [put]
func Create(context *types.Ctx) {
	deviceID := context.Param("device_id")
	if !utils.IsUUID(deviceID) {
		zap.S().Infof("device id %s is not a valid UUID", deviceID)
		context.JSONError(http.StatusBadRequest, errors.New("client_id should be a valid UUID"))
		return
	}

	channel := context.PostForm("channel")
	if !push.IsSenderIdValid(enum.Sender(channel)) {
		zap.S().Infof("channel %s is not supported", channel)
		context.JSONError(http.StatusBadRequest, errors.New("channel is not supported"))
		return
	}

	token := context.PostForm("token")
	deviceName := context.PostForm("device_name")
	deviceMeta := context.PostForm("device_meta")

	_, isChannelChange, oldDevice, ok := dao.Device.EnsureDevice(context,
		deviceID,
		enum.Sender(channel),
		token,
		deviceName,
		deviceMeta,
		context.User,
	)

	if isChannelChange {
		channel, err := push.GetSender(oldDevice.Channel)
		if err != nil {
			zap.S().Errorf("failed to get channel %s", oldDevice.Channel)
			context.JSONError(http.StatusInternalServerError, err)
			return
		}
		if dac, ok := channel.(pushTypes.SenderWithDeviceDeleteAwareness); ok {
			err = dac.OnDeleteDevice(context, oldDevice)
			if err != nil {
				zap.S().Errorf("failed to call channel %s OnDeleteDevice", oldDevice.Channel)
				context.JSONError(http.StatusInternalServerError, err)
				return
			}
		}
	}

	context.JSONResult(ok)
}
