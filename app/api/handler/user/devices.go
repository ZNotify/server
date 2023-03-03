package user

import (
	"errors"
	"net/http"

	"github.com/ZNotify/server/app/api/common"
	"github.com/ZNotify/server/app/db/dao"
)

// Devices
//
//	@Summary      Get user devices
//	@Id           getDevicesByUserSecret
//	@Tags         User
//	@Description  Delete device with device_id
//	@Param        user_secret  path  string  true  "Secret of user"
//	@Produce      json
//	@Success      200  {object}  common.Response[[]models.Device]
//	@Router       /{user_secret}/devices [get]
func Devices(context *common.Context) {
	devices, ok := dao.Device.GetUserDevices(context, context.User)
	if !ok {
		context.JSONError(http.StatusInternalServerError, errors.New("can't get devices"))
		return
	}
	context.JSONResult(devices)
}
