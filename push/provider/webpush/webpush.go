package webpush

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/pkg/errors"

	"notify-api/ent/dao"
	"notify-api/push/item"
	pushTypes "notify-api/push/types"
)

var client = &http.Client{}

type Provider struct {
	VAPIDPublicKey  string
	VAPIDPrivateKey string
	Mailto          string
}

func (p *Provider) Init(cfg pushTypes.Config) error {
	p.VAPIDPublicKey = cfg[VAPIDPublicKey]
	p.VAPIDPrivateKey = cfg[VAPIDPrivateKey]
	p.Mailto = cfg[Mailto]
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

func (p *Provider) Send(msg *item.PushMessage) error {
	tokens, err := dao.DeviceDao.GetUserChannelTokens(msg.UserID, p.Name())
	if err != nil {
		return errors.WithStack(err)
	}
	if len(tokens) == 0 {
		return nil
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	option := p.getOption()
	switch msg.Priority {
	case item.PriorityHigh:
		option.Urgency = webpush.UrgencyHigh
	case item.PriorityNormal:
		option.Urgency = webpush.UrgencyNormal
	case item.PriorityLow:
		option.Urgency = webpush.UrgencyLow
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	size := len(tokens)
	c := make(chan error, 0)
	for _, v := range tokens {
		s := &webpush.Subscription{}
		err = json.Unmarshal([]byte(v), &s)
		if err != nil {
			return err
		}

		go func() {
			_, err := webpush.SendNotificationWithContext(ctx, data, s, option)
			c <- err
		}()
	}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-c:
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

type Config struct {
	VAPIDPublicKey  string
	VAPIDPrivateKey string
	Mailto          string
}

const VAPIDPublicKey = "VAPIDPublicKey"
const VAPIDPrivateKey = "VAPIDPrivateKey"
const Mailto = "Mailto"

func (p *Provider) Config() []string {
	return []string{VAPIDPublicKey, VAPIDPrivateKey, Mailto}
}

func (p *Provider) Name() string {
	return "WebPush"
}
