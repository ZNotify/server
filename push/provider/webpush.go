package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"notify-api/db/model"
	"notify-api/push/types"
	"notify-api/serve/middleware"
	"os"
	"time"
)

var webPushClient = &http.Client{}

type WebPushProvider struct {
	WebPushOption   *webpush.Options
	WebPushClient   *http.Client
	VAPIDPublicKey  string
	VAPIDPrivateKey string
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

	return nil
}

func (p *WebPushProvider) Send(msg *types.Message) error {
	var tokens []string
	subs, err := model.WebSubUtils.Get(msg.UserID)
	if err != nil {
		return err
	}
	for _, v := range subs {
		tokens = append(tokens, v.Subscription)
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	size := len(tokens)
	c := make(chan error, 0)
	for _, v := range tokens {
		s := &webpush.Subscription{}
		err = json.Unmarshal([]byte(v), &s)
		if err != nil {
			return err
		}

		go func() {
			_, err := webpush.SendNotificationWithContext(ctx, data, s, p.WebPushOption)
			c <- err
		}()
	}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-c:
			if err != nil {
				return err
			} else {
				size--
			}
			if size == 0 {
				return nil
			}
		}
	}
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

func (p *WebPushProvider) Name() string {
	return "WebPush"
}

func (p *WebPushProvider) Handler(context *gin.Context) {
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
		context.String(http.StatusOK, "Subscription saved.")
	}
}

func (p *WebPushProvider) HandlerPath() string {
	return "/:user_id/web/sub"
}

func (p *WebPushProvider) HandlerMethod() string {
	return "PUT"
}