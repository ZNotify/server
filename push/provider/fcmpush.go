package provider

import (
	"context"
	"fmt"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/pkg/errors"
	"google.golang.org/api/option"

	"notify-api/db/model"
	"notify-api/push/types"
)

type FCMProvider struct {
	FCMClient     *messaging.Client
	FCMCredential []byte
}

func (p *FCMProvider) Init() error {
	opt := option.WithCredentialsJSON(p.FCMCredential)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}
	p.FCMClient, err = app.Messaging(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (p *FCMProvider) Send(msg *types.Message) error {
	var tokens []string
	tokens, err := model.TokenUtils.GetChannelTokens(msg.UserID, p.Name())

	if err != nil {
		return errors.WithStack(err)
	}

	if len(tokens) == 0 {
		return nil
	}

	// https://firebase.google.com/docs/cloud-messaging/send-message#example-notification-click-action
	fcmMsg := messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: msg.Title,
			Body:  msg.Content,
		},
		Data: map[string]string{
			"user_id":    msg.UserID,
			"long":       msg.Long,
			"msg_id":     msg.ID,
			"title":      msg.Title,
			"content":    msg.Content,
			"created_at": msg.CreatedAt.Format(time.RFC3339),
		},
		Android: &messaging.AndroidConfig{
			Notification: &messaging.AndroidNotification{
				ClickAction: "TranslucentActivity",
			},
		},
		Tokens: tokens,
	}
	_, err = p.FCMClient.SendMulticast(context.Background(), &fcmMsg)
	if err != nil {
		return err
	}
	return nil
}

func (p *FCMProvider) Check(auth types.SenderAuth) error {
	FCMCredential, ok := auth["FCMCredential"]
	if !ok {
		return fmt.Errorf("FCMCredential is not set")
	} else {
		p.FCMCredential = []byte(FCMCredential)
		return nil
	}
}

func (p *FCMProvider) Name() string {
	return "FCM"
}
