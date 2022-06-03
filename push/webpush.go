package push

import (
	"encoding/json"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/ZNotify/server/config"
	"github.com/ZNotify/server/db"
	"github.com/ZNotify/server/db/entity"
	"github.com/ZNotify/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
)

var webPushOption *webpush.Options
var webPushClient *http.Client = &http.Client{}

func SendViaWebPush(msg *entity.Message) error {
	var tokens []entity.WebSubscription
	dbResult := db.DB.Where("user_id = ?", msg.UserID).Find(&tokens)
	if dbResult.Error != nil {
		return dbResult.Error
	}

	var subs []string
	for _, v := range tokens {
		subs = append(subs, v.Subscription)
	}

	if len(subs) == 0 {
		return nil
	}

	data, err := msg.ToJSON()
	if err != nil {
		return err
	}

	for _, v := range subs {
		s := &webpush.Subscription{}
		err = json.Unmarshal([]byte(v), &s)
		if err != nil {
			return err
		}
		_, err := webpush.SendNotification([]byte(data), s, webPushOption)
		if err != nil {
			return err
		}
	}
	defer webPushClient.CloseIdleConnections()

	return nil
}

func InitWebPushOption() {
	webPushOption = &webpush.Options{
		HTTPClient:      webPushClient,
		TTL:             60 * 60 * 24,
		Subscriber:      "zxilly@outlook.com",
		VAPIDPublicKey:  config.VAPIDPublicKey,
		VAPIDPrivateKey: config.VAPIDPrivateKey,
		Urgency:         webpush.UrgencyHigh, // Always send notification, even low battery
	}
}

func SetWebPushSubscription(context *gin.Context) {
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
	db.DB.Model(&entity.WebSubscription{}).
		Where("user_id = ?", userID).
		Where("subscription = ?", tokenString).
		Count(&cnt)
	if cnt > 0 {
		context.String(http.StatusNotModified, "Token already exists")
		return
	} else {
		sub := &entity.WebSubscription{
			ID:           uuid.New().String(),
			UserID:       userID,
			Subscription: tokenString,
		}
		db.DB.Create(sub)
		context.String(http.StatusOK, "Subscription saved.")
	}
}
