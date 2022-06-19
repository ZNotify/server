package providers

import (
	"encoding/json"
	"fmt"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"notify-api/db/entity"
	"notify-api/push"
	"notify-api/serve/middleware"
	"os"
)

var webPushClient = &http.Client{}

type WebPushProvider struct {
	WebPushOption   *webpush.Options
	WebPushClient   *http.Client
	VAPIDPublicKey  string
	VAPIDPrivateKey string
}

func (p *WebPushProvider) Init(e *gin.Engine) error {
	p.WebPushOption = &webpush.Options{
		HTTPClient:      webPushClient,
		TTL:             60 * 60 * 24,
		Subscriber:      "zxilly@outlook.com",
		VAPIDPublicKey:  p.VAPIDPublicKey,
		VAPIDPrivateKey: p.VAPIDPrivateKey,
		Urgency:         webpush.UrgencyHigh, // Always send notification, even low battery
	}
	p.WebPushClient = &http.Client{}

	e.PUT("/:user_id/web/sub", webPushHandler)

	return nil
}

func (p *WebPushProvider) Send(msg *push.Message) error {
	var tokens = entity.WebSubUtils.Get(msg.UserID)

	var subs []string
	for _, v := range tokens {
		subs = append(subs, v.Subscription)
	}

	if len(subs) == 0 {
		return nil
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	for _, v := range subs {
		s := &webpush.Subscription{}
		err = json.Unmarshal([]byte(v), &s)
		if err != nil {
			return err
		}
		_, err := webpush.SendNotification(data, s, p.WebPushOption)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *WebPushProvider) Check() error {
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

func webPushHandler(context *gin.Context) {
	userID := context.GetString(middleware.UserIdKey)

	token, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}
	tokenString := string(token)

	cnt := entity.WebSubUtils.Count(userID, tokenString)
	if cnt > 0 {
		context.String(http.StatusNotModified, "Token already exists")
		return
	} else {
		err := entity.WebSubUtils.Add(userID, tokenString)
		if err != nil {
			context.String(http.StatusInternalServerError, err.Error())
			return
		}
		context.String(http.StatusOK, "Subscription saved.")
	}
}
