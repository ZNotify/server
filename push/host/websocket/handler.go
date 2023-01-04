package websocket

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	dao2 "notify-api/ent/dao"
	"notify-api/push/item"
	"notify-api/server/types"
	"notify-api/utils/ds"
)

func (h *Host) HandlerPath() string {
	return "/host/conn"
}

func (h *Host) HandlerMethod() string {
	return "GET"
}

func (h *Host) Handler(context *types.Ctx) {
	userID := context.UserID

	deviceId := context.GetHeader("X-Device-ID")
	if deviceId == "" {
		zap.S().Infof("user %s connect without device ID", userID)
		context.JSONError(http.StatusBadRequest, errors.New("no device id"))
		return
	}

	device, err := dao2.Device.GetDevice(deviceId)
	if err != nil {
		if err == dao2.ErrNotFound {
			zap.S().Infof("user %s device %s not found", userID, deviceId)
			context.JSONError(http.StatusUnauthorized, errors.New("token not found"))
			return
		} else {
			zap.S().Errorf("get user %s token error: %v", userID, err)
			context.JSONError(http.StatusInternalServerError, errors.WithStack(err))
			return
		}
	}
	if device.Channel != h.Name() {
		zap.S().Infof("device %s channel not match", deviceId)
		context.JSONError(http.StatusUnauthorized, errors.New("device current channel is not WebSocket"))
		return
	}

	// string to uint device.Meta
	msgID, err := strconv.ParseUint(device.Meta, 10, 64)
	if err != nil {
		zap.S().Errorf("device %s meta error: %v", deviceId, err)
		context.JSONError(http.StatusInternalServerError, errors.WithStack(err))
		return
	}

	pendingMessages, err := dao2.Message.GetUserMessageAfter(userID, msgID)
	if err != nil {
		zap.S().Errorf("get user %s message error: %v", userID, err)
		context.JSONError(http.StatusInternalServerError, errors.WithStack(err))
		return
	}

	conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		zap.S().Errorf("upgrade error: %v", err)
		return
	}
	client := &wsClient{
		manager:  h.manager,
		conn:     conn,
		send:     ds.NewUnboundedChan[*item.PushMessage](2),
		userID:   userID,
		deviceID: deviceId,
		once:     sync.Once{},
	}
	h.manager.register <- client

	go client.writeRoutine()
	go client.readRoutine()

	for _, msg := range pendingMessages {
		msg := item.PushMessage{
			MessageID: msg.MessageID,
			UserID:    msg.UserID,
			Title:     msg.Title,
			Content:   msg.Content,
			Long:      msg.Long,
			CreatedAt: msg.CreatedAt,
			Priority:  item.PriorityHigh,
		}
		client.send.In <- &msg
	}
}
