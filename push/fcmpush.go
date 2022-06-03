package push

import (
	"context"
	"firebase.google.com/go/v4/messaging"
	"github.com/ZNotify/server/config"
	"github.com/ZNotify/server/db"
	"github.com/ZNotify/server/db/entity"
	"github.com/ZNotify/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"time"
)

func SendViaFCM(msg *entity.Message) error {
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
			"userID":    msg.UserID,
			"long":      msg.Long,
			"msgID":     msg.ID,
			"title":     msg.Title,
			"content":   msg.Content,
			"createdAt": msg.CreatedAt.Format(time.RFC3339),
		},
		Android: &messaging.AndroidConfig{
			Notification: &messaging.AndroidNotification{
				ClickAction: "TranslucentActivity",
			},
		},
		Tokens: registrationIDs,
	}
	_, err := config.FCMClient.SendMulticast(context.Background(), &fcmMsg)
	if err != nil {
		return err
	}
	return nil
}

func SetFCMToken(context *gin.Context) {
	userID, err := utils.RequireAuth(context)
	if err != nil {
		utils.BreakOnError(context, err)
		return
	}

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
	} else {
		user := entity.FCMTokens{
			ID:             uuid.New().String(),
			UserID:         userID,
			RegistrationID: tokenString,
		}
		db.DB.Create(&user)
		context.String(http.StatusOK, "Registration ID saved.")
	}
}
