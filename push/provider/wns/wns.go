package wns

import (
	"context"

	"notify-api/push/enum"
	pushTypes "notify-api/push/item"
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
