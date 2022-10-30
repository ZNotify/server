package push

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"notify-api/push/host"
	"notify-api/push/provider"
	pushTypes "notify-api/push/types"
	serveTypes "notify-api/serve/types"
	"notify-api/utils"
)

type senders struct {
	senders []pushTypes.Sender
}

var Senders = senders{
	senders: []pushTypes.Sender{
		new(provider.FCMProvider),
		new(provider.WebPushProvider),
		new(host.WebSocketHost),
	},
}

func (p *senders) Has(channel string) bool {
	for _, v := range p.senders {
		if v.Name() == channel {
			return true
		}
	}
	return false
}

func (p *senders) Send(msg *pushTypes.Message) error {
	if utils.IsTestInstance() {
		return nil
	}

	var errs []string
	var wg sync.WaitGroup
	wg.Add(len(p.senders))
	for _, v := range p.senders {
		go func(sender pushTypes.Sender) {
			defer wg.Done()
			// log.Println("Sending message to", sender.Name())
			pe := sender.Send(msg)
			if pe != nil {
				errString := fmt.Sprintf("Send message to %s failed: %v", sender.Name(), pe)
				errs = append(errs, errString)
			}
		}(v)
	}
	if utils.WaitTimeout(&wg, 5*time.Second) {
		return errors.New("send timeout")
	}
	if len(errs) > 0 {
		val := ""
		for _, v := range errs {
			val += v + "\n"
		}
		return errors.New(val)
	}
	return nil
}

func (p *senders) Init() {
	for _, sender := range p.senders {
		if pv, ok := sender.(pushTypes.Provider); ok {
			if utils.IsTestInstance() {
				continue
			}

			if err := pv.Check(); err != nil {
				log.Fatalf("Provider %s check failed: %s", pv.Name(), err)
			}
			err := pv.Init()
			if err != nil {
				log.Fatalf("Provider %s init failed: %s", pv.Name(), err)
				return
			} else {
				log.Printf("Provider %s init success", pv.Name())
			}
		}

		if hv, ok := sender.(pushTypes.Host); ok {
			if err := hv.Init(); err != nil {
				log.Fatalf("Host %s init failed: %s", hv.Name(), err)
				return
			} else {
				log.Printf("Host %s init success", hv.Name())
			}

			if err := hv.Start(); err != nil {
				log.Fatalf("Host %s start failed: %s", hv.Name(), err)
				return
			} else {
				log.Printf("Host %s start success", hv.Name())
			}
		}
	}
}

func (p *senders) RegisterRouter(e *gin.RouterGroup) error {
	if len(p.senders) == 0 && !utils.IsTestInstance() {
		return errors.New("no sender found")
	}
	for _, v := range p.senders {
		if pv, ok := v.(pushTypes.SenderWithHandler); ok {
			e.Handle(pv.HandlerMethod(), pv.HandlerPath(), serveTypes.WrapHandler(pv.Handler))
		}
	}
	return nil
}
