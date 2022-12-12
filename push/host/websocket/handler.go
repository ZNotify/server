package websocket

import (
	"net/http"
	"sync"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"notify-api/db/util"
	"notify-api/push/entity"
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

	device, err := util.DeviceUtil.GetDevice(deviceId)
	if err != nil {
		if err == util.ErrNotFound {
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

	sinceTime, err := time.Parse(time.RFC3339Nano, device.Meta)
	if err != nil {
		zap.S().Infof("parse time error: %v", err)
		context.JSONError(http.StatusBadRequest, errors.WithStack(err))
		return
	}
	// 2022-09-18T11:14:00+08:00 as zero point

	pendingMessages, err := util.MessageUtil.GetUserMessageAfter(userID, sinceTime)
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
		send:     ds.NewUnboundedChan[*entity.PushMessage](2),
		userID:   userID,
		deviceID: deviceId,
		once:     sync.Once{},
	}
	h.manager.register <- client

	go client.writeRoutine()
	go client.readRoutine()

	for _, msg := range pendingMessages {
		msg := entity.PushMessage{
			MessageID: msg.MessageID,
			UserID:    msg.UserID,
			Title:     msg.Title,
			Content:   msg.Content,
			Long:      msg.Long,
			CreatedAt: msg.CreatedAt,
			Priority:  entity.PriorityHigh,
		}
		client.send.In <- &msg
	}
}
