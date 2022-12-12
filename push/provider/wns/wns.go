package wns

import (
	pushTypes "notify-api/push/entity"
)

type Provider struct {
}

func (p *Provider) Send(msg *pushTypes.PushMessage) error {
	//TODO implement me
	panic("implement me")
}

func (p *Provider) Init() error {
	//TODO implement me
	panic("implement me")
}

func (p *Provider) Name() string {
	// Windows Push Notification Services
	return "WNS"
}
