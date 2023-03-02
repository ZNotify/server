package webpush

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/SherClockHolmes/webpush-go"

	"notify-api/app/api/models"
	"notify-api/app/config/sender"
	"notify-api/app/db/dao"
	"notify-api/app/manager/push/enum"
	"notify-api/app/manager/push/interfaces"
	"notify-api/app/manager/push/item"
)

var client = &http.Client{}

type Provider struct {
	VAPIDPublicKey  string
	VAPIDPrivateKey string
	Mailto          string
}

func (p *Provider) Init(cfg interfaces.Config) error {
	config, ok := cfg.(senderConfig.WebPushConfig)
	if !ok {
		return errors.New("webpush config type error")
	}
	p.VAPIDPublicKey = config.VAPIDPublicKey
	p.VAPIDPrivateKey = config.VAPIDPrivateKey
	p.Mailto = config.Mailto
	return nil
}

func (p *Provider) getOption() *webpush.Options {
	return &webpush.Options{
		HTTPClient:      client,
		TTL:             60 * 60 * 24,
		Subscriber:      p.Mailto,
		VAPIDPublicKey:  p.VAPIDPublicKey,
		VAPIDPrivateKey: p.VAPIDPrivateKey,
	}
}

func (p *Provider) Send(ctx context.Context, msg *item.PushMessage) error {
	tokens, ok := dao.Device.GetUserDeviceChannelTokens(ctx, msg.User, p.Name())
	if !ok {
		return errors.New("webpush get user device channel tokens failed")
	}
	if len(tokens) == 0 {
		return nil
	}

	data, err := json.Marshal(models.FromPushMessage(*msg))
	if err != nil {
		return err
	}

	option := p.getOption()
	switch msg.Priority {
	case enum.PriorityHigh:
		option.Urgency = webpush.UrgencyHigh
	case enum.PriorityNormal:
		option.Urgency = webpush.UrgencyNormal
	case enum.PriorityLow:
		option.Urgency = webpush.UrgencyLow
	}

	size := len(tokens)
	c := make(chan error)
	for _, v := range tokens {
		s := &webpush.Subscription{}
		err = json.Unmarshal([]byte(v), &s)
		if err != nil {
			return err
		}

		go func() {
			resp, err := webpush.SendNotificationWithContext(ctx, data, s, option)
			_ = resp.Body.Close()
			c <- err
		}()
	}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err = <-c:
			if err != nil {
				return err
			} else {
				size--
			}
			if size == 0 {
				return nil
			}
		}
	}
}

func (p *Provider) Name() enum.Sender {
	return enum.SenderWebPush
}
