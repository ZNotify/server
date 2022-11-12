package websocket

import (
	"net/http"
	"sync"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"notify-api/db/model"
	pushTypes "notify-api/push/types"
	"notify-api/serve/types"
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

	token, err := model.TokenUtils.GetUserDeviceToken(userID, deviceId)
	if err != nil {
		if err == model.ErrNotFound {
			zap.S().Infof("user %s device %s token not found", userID, deviceId)
			context.JSONError(http.StatusUnauthorized, errors.New("token not found"))
			return
		} else {
			zap.S().Errorf("get user %s token error: %v", userID, err)
			context.JSONError(http.StatusInternalServerError, errors.WithStack(err))
			return
		}
	}
	if token.Channel != h.Name() {
		zap.S().Infof("user %s channel not match", userID)
		context.JSONError(http.StatusUnauthorized, errors.New("device current channel is not WebSocket"))
		return
	}

	sinceTime, err := time.Parse(time.RFC3339Nano, token.TokenMeta)
	if err != nil {
		zap.S().Infof("parse time error: %v", err)
		context.JSONError(http.StatusBadRequest, errors.WithStack(err))
		return
	}
	// 2022-09-18T11:14:00+08:00 as zero point

	pendingMessages, err := model.MessageUtils.GetUserMessageAfter(userID, sinceTime)
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
		send:     ds.NewUnboundedChan[*pushTypes.Message](2),
		userID:   userID,
		deviceID: deviceId,
		once:     sync.Once{},
	}
	h.manager.register <- client

	go client.writeRoutine()
	go client.readRoutine()

	for _, msg := range pendingMessages {
		msg := pushTypes.Message{
			ID:        msg.ID,
			UserID:    msg.UserID,
			Title:     msg.Title,
			Content:   msg.Content,
			Long:      msg.Long,
			CreatedAt: msg.CreatedAt,
			Priority:  pushTypes.PriorityHigh,
		}
		client.send.In <- &msg
	}
}
