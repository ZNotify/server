package types

import (
	"context"

	"notify-api/ent/generate"
	"notify-api/push/enum"
	"notify-api/push/item"
	"notify-api/server/types"
)

type Config = map[string]string

type Sender interface {
	Send(ctx context.Context, msg *item.PushMessage) error
	Name() enum.Sender
}

type SenderWithoutConfig interface {
	Sender
	Init() error
}

type SenderWithConfig interface {
	Sender
	Init(config Config) error
	Config() []string
}

type SenderWithBackground interface {
	Sender
	Setup() error
}

type SenderWithHandler interface {
	Sender
	Handler(ctx *types.Ctx)
	HandlerPath() string
	HandlerMethod() string
}

type SenderWithDeviceDeleteAwareness interface {
	Sender
	OnDeleteDevice(ctx *types.Ctx, device *generate.Device) error
}
