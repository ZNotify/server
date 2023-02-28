package interfaces

import (
	"context"

	"notify-api/app/common"
	"notify-api/app/manager/push/enum"
	"notify-api/app/manager/push/item"
	"notify-api/db/ent/generate"
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
	Handler(ctx *common.Context)
	HandlerPath() string
	HandlerMethod() string
}

type SenderWithDeviceDeleteAwareness interface {
	Sender
	OnDeleteDevice(ctx *common.Context, device *generate.Device) error
}
