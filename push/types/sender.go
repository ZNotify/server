package types

import (
	"context"

	"notify-api/push/item"
	"notify-api/server/types"
)

type Config = map[string]string

type Sender interface {
	Send(ctx context.Context, msg *item.PushMessage) error
	Name() string
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

type Host interface {
	Sender
	Start() error
}

type SenderWithHandler interface {
	Sender
	Handler(ctx *types.Ctx)
	HandlerPath() string
	HandlerMethod() string
}
