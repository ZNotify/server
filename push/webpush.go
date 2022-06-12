package push

import (
	"encoding/json"
	"fmt"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/ZNotify/server/db"
	"github.com/ZNotify/server/db/entity"
	"github.com/ZNotify/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"os"
)

var webPushClient = &http.Client{}

type WebPushProvider struct {
	WebPushOption   *webpush.Options
	WebPushClient   *http.Client
	VAPIDPublicKey  string
	VAPIDPrivateKey string
}

func (p *WebPushProvider) init() (Provider, error) {
	err := p.check()
	if err != nil {
		return nil, err
	}
	p.WebPushOption = &webpush.Options{
		HTTPClient:      webPushClient,
		TTL:             60 * 60 * 24,
		Subscriber:      "zxilly@outlook.com",
		VAPIDPublicKey:  p.VAPIDPublicKey,
		VAPIDPrivateKey: p.VAPIDPrivateKey,
		Urgency:         webpush.UrgencyHigh, // Always send notification, even low battery
	}
	p.WebPushClient = &http.Client{}
	return p, nil
}

func (p *WebPushProvider) send(msg *entity.Message) error {
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
		_, err := webpush.SendNotification([]byte(data), s, p.WebPushOption)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *WebPushProvider) check() error {
	VAPIDPublicKey := os.Getenv("VAPIDPublicKey")
	VAPIDPrivateKey := os.Getenv("VAPIDPrivateKey")
	if VAPIDPublicKey == "" || VAPIDPrivateKey == "" {
		return fmt.Errorf("VAPIDPublicKey or VAPIDPrivateKey is empty")
	} else {
		p.VAPIDPublicKey = VAPIDPublicKey
		p.VAPIDPrivateKey = VAPIDPrivateKey
		return nil
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
