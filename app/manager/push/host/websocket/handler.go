package websocket

import (
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"notify-api/app/api/common"
	"notify-api/app/db/dao"
	"notify-api/app/db/helper"
	"notify-api/app/manager/push/item"
)

func (h *Host) HandlerPath() string {
	return "/conn"
}

func (h *Host) HandlerMethod() string {
	return http.MethodGet
}

// Handler
//
//	@Summary      Endpoint for websocket connection
//	@Id           websocket
//	@Tags         Push
//	@Description  Work as a fallback strategy for device without public push provider, each frame in this connection will be a push message
//	@Param        X-Device-ID  header  string  true  "Device ID, usually a UUID"
//	@Param        user_secret  path    string  true  "Secret of user"
//	@Produce      json
//	@Success      200  {object}  models.Message
//	@Failure      400  {object}  common.BadRequestResponse
//	@Failure      401  {object}  common.UnauthorizedResponse
//	@Router       /{user_secret}/conn [get]
func (h *Host) Handler(context *common.Context) {
	deviceId := context.GetHeader("X-Device-ID")
	if deviceId == "" {
		zap.S().Infof("user %s connect without device ID", helper.GetReadableName(context.User))
		context.JSONError(http.StatusBadRequest, errors.New("no device id"))
		return
	}

	device, ok := dao.Device.GetUserDeviceByIdentifier(context, context.User, deviceId)
	if !ok {
		zap.S().Infof("user %s connect with invalid device ID", helper.GetReadableName(context.User))
		context.JSONError(http.StatusBadRequest, errors.New("invalid device id"))
		return
	}

	var pendingMessages []*item.PushMessage
	if device.DeviceMeta == "" {
		pendingMessages = make([]*item.PushMessage, 0)
		ok := dao.Device.UpdateDeviceChannelMeta(context, device, strconv.FormatInt(dao.SequenceID.Load(), 10))
		if !ok {
			zap.S().Errorf("failed to update device channel meta")
			context.JSONError(http.StatusInternalServerError, errors.New("failed to update device channel meta"))
			return
		}
	} else {
		msgID, err := strconv.ParseInt(device.DeviceMeta, 10, 64)
		if err != nil {
			zap.S().Errorf("failed to parse device meta %s", device.DeviceMeta)
			context.JSONError(http.StatusInternalServerError, err)
			return
		}
		modelMessages, ok := dao.Message.GetUserMessagesAfterID(context, context.User, int(msgID))
		if !ok {
			zap.S().Errorf("failed to get user messages after id %d", msgID)
			context.JSONError(http.StatusInternalServerError, err)
			return
		}
		pendingMessages = make([]*item.PushMessage, len(modelMessages))
		for i, msg := range modelMessages {
			pendingMessages[i] = item.FromModelMessageWithUser(msg, context.User)
		}
	}

	conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		zap.S().Errorf("upgrade error: %v", err)
		return
	}

	client := newClient(context, device, conn, manager.unregister)
	manager.register(client)
	client.run()

	for _, msg := range pendingMessages {
		client.send <- msg
	}
}
