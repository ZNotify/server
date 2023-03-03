package device

import (
	"github.com/ZNotify/server/app/api/common"
	"github.com/ZNotify/server/app/db/dao"
)

// Delete godoc
//
//	@Summary      Delete device
//	@Id           deleteDevice
//	@Tags         Device
//	@Description  Delete device with device_id
//	@Param        user_secret  path  string  true  "Secret of user"
//	@Param        device_id    path  string  true  "The identifier of device, should be a UUID"
//	@Produce      json
//	@Success      200  {object}  common.Response[bool]
//	@Router       /{user_secret}/device/{device_id} [delete]
func Delete(context *common.Context) {
	deviceId := context.Param("device_id")
	ok := dao.Device.DeleteDeviceByIdentifier(context, deviceId)
	context.JSONResult(ok)
}
