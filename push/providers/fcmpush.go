package providers

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"io"
	"net/http"
	"notify-api/db/model"
	"notify-api/push"
	"notify-api/serve/middleware"
	"notify-api/user"
	"os"
	"time"
)

type FCMProvider struct {
	FCMClient     *messaging.Client
	FCMCredential []byte

	regIDCache map[string][]string
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

	users := user.Controller.Users()
	p.regIDCache = make(map[string][]string)
	for _, u := range users {
		regs, err := model.FCMTokenUtils.Get(u)
		if err != nil {
			return err
		}
		tokens := make([]string, 0)
		for _, r := range regs {
			tokens = append(tokens, r.RegistrationID)
		}
		p.regIDCache[u] = tokens
	}
	return nil
}

func (p *FCMProvider) Send(msg *push.Message) error {

	if len(p.regIDCache[msg.UserID]) == 0 {
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
		Tokens: p.regIDCache[msg.UserID],
	}
	_, err := p.FCMClient.SendMulticast(context.Background(), &fcmMsg)
	if err != nil {
		return err
	}
	return nil
}

func (p *FCMProvider) Check() error {
	FCMCredential := []byte(os.Getenv("FCMCredential"))
	if len(FCMCredential) == 0 {
		return fmt.Errorf("FCMCredential is not set")
	} else {
		p.FCMCredential = FCMCredential
		return nil
	}
}

func (p *FCMProvider) ChannelName() string {
	return "FCM"
}

func (p *FCMProvider) ProviderHandler(context *gin.Context) {
	userID := context.GetString(middleware.UserIdKey)

	token, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}
	tokenString := string(token)

	cnt, err := model.FCMTokenUtils.Count(userID, tokenString)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	// TODO: update user with same token
	if cnt > 0 {
		context.String(http.StatusNotModified, "Token already exists")
		return
	} else {
		_, err := model.FCMTokenUtils.Add(userID, tokenString)
		if err != nil {
			context.String(http.StatusInternalServerError, err.Error())
			return
		}
		p.regIDCache[userID] = append(p.regIDCache[userID], tokenString)
		context.String(http.StatusOK, "Registration ID saved.")
		return
	}
}

func (p *FCMProvider) ProviderHandlerPath() string {
	return "/:user_id/fcm/token"
}

func (p *FCMProvider) ProviderHandlerMethod() string {
	return "PUT"
}
