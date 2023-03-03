package interfaces

import (
	"context"

	"github.com/ZNotify/server/app/api/common"
	"github.com/ZNotify/server/app/db/ent/generate"
	"github.com/ZNotify/server/app/manager/push/enum"
	"github.com/ZNotify/server/app/manager/push/item"
)

type Config any

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
