package controller

import (
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"notify-api/ent/dao"
	"notify-api/push"
	"notify-api/server/types"
	"notify-api/utils"
)

// Token godoc
//
//	@Summary		Create or update token
//	@Description	Create or update token of device
//	@Param			user_id		path		string	true	"user_id"
//	@Param			client_id	path		string	true	"client_id should be a valid UUID"
//	@Param			channel		formData	string	true	"channel can be used. Sometimes less than document."	Enums(TelegramHost, WebSocketHost, FCM, WebPush, WNS)
//	@Param			token		formData	string	false	"token"
//	@Param			info		formData	string	false	"Additional info about client"
//	@Produce		json
//	@Success		200	{object}	types.Response[bool]
//	@Failure		400	{object}	types.BadRequestResponse
//	@Failure		401	{object}	types.UnauthorizedResponse
//	@Router			/{user_id}/token/{client_id} [put]
func Token(context *types.Ctx) {
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

	info := context.PostForm("info")
	if info == "" {
		info = context.Request.UserAgent()
	}

	err := dao.DeviceDao.CreateOrUpdate(context.UserID, deviceID, channel, token, info)
	if err != nil {
		zap.S().Errorw("create or update token error", "error", err)
		context.JSONError(http.StatusInternalServerError, errors.WithStack(err))
		return
	}

	tokenMeta, ok := push.TryGetInitialDeviceMeta(channel)
	if ok {
		err = dao.DeviceDao.UpdateDeviceMeta(deviceID, tokenMeta)
		if err != nil {
			zap.S().Errorw("update token meta error", "error", err)
			context.JSONError(http.StatusInternalServerError, errors.WithStack(err))
			return
		}
	}

	context.JSONResult(true)
}
