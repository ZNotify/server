package fcm

import (
	"context"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/pkg/errors"
	"google.golang.org/api/option"

	"notify-api/ent/dao"
	"notify-api/push/enum"
	"notify-api/push/item"
	"notify-api/push/types"
)

type Provider struct {
	Client     *messaging.Client
	Credential []byte
}

func (p *Provider) Init(cfg types.Config) error {
	p.Credential = []byte(cfg[Credential])

	opt := option.WithCredentialsJSON(p.Credential)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}
	p.Client, err = app.Messaging(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (p *Provider) Send(ctx context.Context, msg *item.PushMessage) error {
	tokens, ok := dao.Device.GetUserDeviceChannelTokens(ctx, msg.User, p.Name())

	if !ok {
		return errors.New("fcm get user device channel tokens failed")
	}

	if len(tokens) == 0 {
		return nil
	}

	var fcmPriority string
	if msg.Priority == enum.PriorityHigh {
		fcmPriority = "high"
	} else {
		fcmPriority = "normal"
	}

	// https://firebase.google.com/docs/cloud-messaging/send-message#example-notification-click-action
	fcmMsg := messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: msg.Title,
			Body:  msg.Content,
		},
		Data: map[string]string{
			"long":       msg.Long,
			"msg_id":     msg.ID.String(),
			"title":      msg.Title,
			"content":    msg.Content,
			"created_at": msg.CreatedAt.Format(time.RFC3339),
		},
		Android: &messaging.AndroidConfig{
			Notification: &messaging.AndroidNotification{
				ClickAction: "TranslucentActivity",
			},
			Priority: fcmPriority,
		},
		Tokens: tokens,
	}
	_, err := p.Client.SendMulticast(ctx, &fcmMsg)
	return err
}

const Credential = "Credential"

func (p *Provider) Config() []string {
	return []string{Credential}
}

func (p *Provider) Name() enum.Sender {
	return enum.SenderFcm
}
