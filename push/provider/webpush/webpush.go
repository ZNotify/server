package webpush

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/pkg/errors"

	"notify-api/db/model"
	pushTypes "notify-api/push/types"
	"notify-api/utils"
	"notify-api/utils/config"
)

var webPushClient = &http.Client{}

type Provider struct {
	VAPIDPublicKey  string
	VAPIDPrivateKey string
	Mailto          string
}

func (p *Provider) Init() error {
	cfg := config.Config.Senders[p.Name()].(Config)
	p.VAPIDPublicKey = utils.TokenClean(cfg.VAPIDPublicKey)
	p.VAPIDPrivateKey = utils.TokenClean(cfg.VAPIDPrivateKey)
	p.Mailto = utils.TokenClean(cfg.Mailto)
	return nil
}

func (p *Provider) getOption() *webpush.Options {
	return &webpush.Options{
		HTTPClient:      webPushClient,
		TTL:             60 * 60 * 24,
		Subscriber:      p.Mailto,
		VAPIDPublicKey:  p.VAPIDPublicKey,
		VAPIDPrivateKey: p.VAPIDPrivateKey,
	}
}

func (p *Provider) Send(msg *pushTypes.Message) error {
	var tokens []string
	tokens, err := model.TokenUtils.GetUserChannelTokens(msg.UserID, p.Name())
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
	case pushTypes.PriorityHigh:
		option.Urgency = webpush.UrgencyHigh
	case pushTypes.PriorityNormal:
		option.Urgency = webpush.UrgencyNormal
	case pushTypes.PriorityLow:
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

func (p *Provider) Config() any {
	return Config{}
}

func (p *Provider) Name() string {
	return "WebPush"
}
