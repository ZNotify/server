package push

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/api/option"
	"io/ioutil"
	"net/http"
	"notify-api/db"
	"notify-api/db/entity"
	"notify-api/serve/middleware"
	"os"
	"time"
)

type FCMProvider struct {
	FCMClient     *messaging.Client
	FCMCredential []byte
}

func (p *FCMProvider) init(e *gin.Engine) error {
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

func (p *FCMProvider) send(msg *entity.Message) error {
	var tokens []entity.FCMTokens
	dbResult := db.DB.Where("user_id = ?", msg.UserID).Find(&tokens)
	if dbResult.Error != nil {
		return dbResult.Error
	}

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

func (p *FCMProvider) check() error {
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

	var cnt int64
	db.DB.Model(&entity.FCMTokens{}).
		Where("user_id = ?", userID).
		Where("registration_id = ?", tokenString).
		Count(&cnt)
	// TODO: update user with same token
	if cnt > 0 {
		context.String(http.StatusNotModified, "Token already exists")
		return
	} else {
		userEntity := entity.FCMTokens{
			ID:             uuid.New().String(),
			UserID:         userID,
			RegistrationID: tokenString,
		}
		db.DB.Create(&userEntity)
		context.String(http.StatusOK, "Registration ID saved.")
		return
	}
}
