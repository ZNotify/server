package providers

import (
	"encoding/json"
	"fmt"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"notify-api/db/model"
	"notify-api/push"
	"notify-api/serve/middleware"
	"notify-api/user"
	"os"
)

var webPushClient = &http.Client{}

type WebPushProvider struct {
	WebPushOption   *webpush.Options
	WebPushClient   *http.Client
	VAPIDPublicKey  string
	VAPIDPrivateKey string

	subsCache map[string][]string
}

func (p *WebPushProvider) Init() error {
	p.WebPushOption = &webpush.Options{
		HTTPClient:      webPushClient,
		TTL:             60 * 60 * 24,
		Subscriber:      "zxilly@outlook.com",
		VAPIDPublicKey:  p.VAPIDPublicKey,
		VAPIDPrivateKey: p.VAPIDPrivateKey,
		Urgency:         webpush.UrgencyHigh, // Always send notification, even low battery
	}
	p.WebPushClient = &http.Client{}

	users := user.Controller.Users()
	p.subsCache = make(map[string][]string)
	for _, v := range users {
		tokens, err := model.WebSubUtils.Get(v)
		if err != nil {
			return err
		}
		var subs []string
		for _, v := range tokens {
			subs = append(subs, v.Subscription)
		}
		p.subsCache[v] = subs
	}

	return nil
}

func (p *WebPushProvider) Send(msg *push.Message) error {
	if len(p.subsCache[msg.UserID]) == 0 {
		return nil
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	for _, v := range p.subsCache[msg.UserID] {
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

func (p *WebPushProvider) ChannelName() string {
	return "WebPush"
}

func (p *WebPushProvider) ProviderHandler(context *gin.Context) {
	userID := context.GetString(middleware.UserIdKey)

	token, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}
	tokenString := string(token)

	cnt := model.WebSubUtils.Count(userID, tokenString)
	if cnt > 0 {
		context.String(http.StatusNotModified, "Token already exists")
		return
	} else {
		_, err := model.WebSubUtils.Add(userID, tokenString)
		if err != nil {
			context.String(http.StatusInternalServerError, err.Error())
			return
		}
		p.subsCache[userID] = append(p.subsCache[userID], tokenString)
		context.String(http.StatusOK, "Subscription saved.")
	}
}

func (p *WebPushProvider) ProviderHandlerPath() string {
	return "/:user_id/web/sub"
}

func (p *WebPushProvider) ProviderHandlerMethod() string {
	return "PUT"
}
