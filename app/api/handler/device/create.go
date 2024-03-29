package device

import (
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/ZNotify/server/app/api/common"
	"github.com/ZNotify/server/app/db/dao"
	"github.com/ZNotify/server/app/manager/push"
	"github.com/ZNotify/server/app/manager/push/enum"
	"github.com/ZNotify/server/app/manager/push/interfaces"
	"github.com/ZNotify/server/app/utils"
)

// Create godoc
//
//	@Summary      Create or update device
//	@Id           createDevice
//	@Tags         Device
//	@Description  Create or update device information
//	@Param        user_secret  path      string       true   "Secret of user"
//	@Param        device_id    path      string       true   "device_id should be a valid UUID"
//	@Param        channel      formData  enum.Sender  true   "channel can be used."
//	@Param        device_name  formData  string       false  "device name"
//	@Param        device_meta  formData  string       false  "additional device meta"
//	@Param        token        formData  string       false  "channel token"
//	@Produce      json
//	@Success      200  {object}  common.Response[bool]
//	@Failure      400  {object}  common.BadRequestResponse
//	@Failure      401  {object}  common.UnauthorizedResponse
//	@Router       /{user_secret}/device/{device_id} [put]
func Create(context *common.Context) {
	deviceID := context.Param("device_id")
	if !utils.IsUUID(deviceID) {
		zap.S().Infof("device id %s is not a valid UUID", deviceID)
		context.JSONError(http.StatusBadRequest, errors.New("client_id should be a valid UUID"))
		return
	}

	channel := context.PostForm("channel")
	if !push.IsSenderActive(enum.Sender(channel)) {
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
		if dac, ok := channel.(interfaces.SenderWithDeviceDeleteAwareness); ok {
			err = dac.OnDeleteDevice(context, oldDevice)
			if err != nil {
				zap.S().Errorf("failed to call channel %s OnDeleteDevice", oldDevice.Channel)
				context.JSONError(http.StatusInternalServerError, err)
				return
			}
		}
	}

	if !ok {
		context.JSONError(http.StatusInternalServerError, errors.New("can't create device"))
		return
	}

	context.JSONResult(true)
}
