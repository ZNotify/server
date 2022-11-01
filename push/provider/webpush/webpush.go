package webpush

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/pkg/errors"

	"notify-api/db/model"
	"notify-api/push/types"
)

var webPushClient = &http.Client{}

type Provider struct {
	WebPushOption   *webpush.Options
	WebPushClient   *http.Client
	VAPIDPublicKey  string
	VAPIDPrivateKey string
}

func (p *Provider) Init() error {
	p.WebPushOption = &webpush.Options{
		HTTPClient:      webPushClient,
		TTL:             60 * 60 * 24,
		Subscriber:      "zxilly@outlook.com",
		VAPIDPublicKey:  p.VAPIDPublicKey,
		VAPIDPrivateKey: p.VAPIDPrivateKey,
		Urgency:         webpush.UrgencyHigh, // Always send notification, even low battery
	}
	p.WebPushClient = &http.Client{}

	return nil
}

func (p *Provider) Send(msg *types.Message) error {
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
			_, err := webpush.SendNotificationWithContext(ctx, data, s, p.WebPushOption)
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

func (p *Provider) Check(auth types.SenderAuth) error {
	VAPIDPublicKey, ok := auth["VAPIDPublicKey"]
	if !ok {
		return fmt.Errorf("VAPIDPublicKey not found")
	}

	VAPIDPrivateKey, ok := auth["VAPIDPrivateKey"]
	if !ok {
		return fmt.Errorf("VAPIDPrivateKey not found")
	}

	p.VAPIDPublicKey = VAPIDPublicKey
	p.VAPIDPrivateKey = VAPIDPrivateKey
	return nil
}

func (p *Provider) Name() string {
	return "WebPush"
}
