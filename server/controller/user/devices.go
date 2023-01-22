package user

import (
	"errors"
	"net/http"

	"notify-api/ent/dao"
	"notify-api/server/types"
)

// Devices
//
//	@Summary		Get user devices
//	@Id				getDevicesByUserSecret
//	@Tags			User
//	@Description	Delete device with device_id
//	@Param			user_secret	path	string	true	"Secret of user"
//	@Produce		json
//	@Success		200	{object}	types.Response[[]entity.Device]
//	@Router			/{user_secret}/devices [get]
func Devices(context *types.Ctx) {
	devices, ok := dao.Device.GetUserDevices(context, context.User)
	if !ok {
		context.JSONError(http.StatusInternalServerError, errors.New("can't get devices"))
		return
	}
	context.JSONResult(devices)
}
