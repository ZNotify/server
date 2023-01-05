package device

import (
	"notify-api/ent/dao"
	"notify-api/server/types"
)

// Delete godoc
//
//	@Summary		Delete device
//	@Description	Delete device with device_id
//	@Param			user_secret	path	string	true	"Secret of user"
//	@Param			device_id	path	string	true	"device_id"
//	@Produce		json
//	@Success		200	{object}	types.Response[bool]
//	@Router			/{user_secret}/device/{device_id} [delete]
func Delete(context *types.Ctx) {
	deviceId := context.Param("device_id")
	ok := dao.Device.DeleteDeviceByIdentifier(context, deviceId)
	context.JSONResult(ok)
}
