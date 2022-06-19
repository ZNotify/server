package providers

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"io/ioutil"
	"net/http"
	"notify-api/db/entity"
	"notify-api/push"
	"notify-api/serve/middleware"
	"os"
	"time"
)

type FCMProvider struct {
	FCMClient     *messaging.Client
	FCMCredential []byte
}

func (p *FCMProvider) Init(e *gin.Engine) error {
	opt := option.WithCredentialsJSON(p.FCMCredential)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}
	p.FCMClient, err = app.Messaging(context.Background())
	if err != nil {
		return err
	}

	e.PUT("/:user_id/fcm/token", fcmTokenHandler)

	return nil
}

func (p *FCMProvider) Send(msg *push.Message) error {
	var tokens = entity.FCMUtils.Get(msg.UserID)

	var registrationIDs []string
	for i := range tokens {
		registrationIDs = append(registrationIDs, tokens[i].RegistrationID)
	}

	if len(registrationIDs) == 0 {
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
		Tokens: registrationIDs,
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

func fcmTokenHandler(context *gin.Context) {
	userID := context.GetString(middleware.UserIdKey)

	token, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}
	tokenString := string(token)

	cnt, err := entity.FCMUtils.Count(userID, tokenString)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	// TODO: update user with same token
	if cnt > 0 {
		context.String(http.StatusNotModified, "Token already exists")
		return
	} else {
		_, err := entity.FCMUtils.Add(userID, tokenString)
		if err != nil {
			context.String(http.StatusInternalServerError, err.Error())
			return
		}
		context.String(http.StatusOK, "Registration ID saved.")
		return
	}
}
