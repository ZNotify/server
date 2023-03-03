package wns

import (
	"context"

	"github.com/ZNotify/server/app/manager/push/enum"
	pushTypes "github.com/ZNotify/server/app/manager/push/item"
)

type Provider struct {
}

func (p *Provider) Send(ctx context.Context, msg *pushTypes.PushMessage) error {
	//TODO implement me
	panic("implement me")
}

func (p *Provider) Init() error {
	//TODO implement me
	panic("implement me")
}

func (p *Provider) Name() enum.Sender {
	// Windows Push Notification Services
	return enum.SenderWns
}
