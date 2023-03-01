package websocket

import (
	"context"

	"notify-api/app/api/common"
	"notify-api/app/db/ent/generate"
	"notify-api/app/manager/push/enum"
	"notify-api/app/manager/push/item"
)

type Host struct {
	manager *clientManager
}

func (h *Host) OnDeleteDevice(ctx *common.Context, device *generate.Device) error {
	manager.unregisterByDevice(device)
	return nil
}

func (h *Host) Send(ctx context.Context, msg *item.PushMessage) error {
	h.manager.send(msg)
	return nil
}

func (h *Host) Name() enum.Sender {
	return enum.SenderWebSocket
}
