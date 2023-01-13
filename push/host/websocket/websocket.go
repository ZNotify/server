package websocket

import (
	"context"

	"notify-api/ent/generate"
	"notify-api/push/enum"
	"notify-api/push/item"
	"notify-api/server/types"
)

type Host struct {
	manager *clientManager
}

func (h *Host) OnDeleteDevice(ctx *types.Ctx, device *generate.Device) error {
	manager.unregisterByDevice(device)
	return nil
}

func (h *Host) Send(ctx context.Context, msg *item.PushMessage) error {
	h.manager.send(msg)
	return nil
}

func (h *Host) Name() string {
	return string(enum.SenderWebSocket)
}
