package websocket

import (
	"context"

	"github.com/ZNotify/server/app/api/common"
	"github.com/ZNotify/server/app/db/ent/generate"
	"github.com/ZNotify/server/app/manager/push/enum"
	"github.com/ZNotify/server/app/manager/push/item"
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
