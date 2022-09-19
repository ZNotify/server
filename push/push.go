package push

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"notify-api/push/host"
	"notify-api/push/provider"
	"notify-api/push/types"
	"notify-api/serve/middleware"
	"notify-api/utils"
	"sync"
)

type senders struct {
	senders []types.Sender
}

var Senders = senders{
	senders: []types.Sender{
		new(provider.FCMProvider),
		new(provider.WebPushProvider),
		new(host.WebSocketHost),
	},
}

func (p *senders) Send(msg *types.Message) error {
	if utils.IsTestInstance() {
		return nil
	}

	var errs []error
	var wg sync.WaitGroup
	wg.Add(len(p.senders))
	for _, v := range p.senders {
		go func(sender *types.Sender) {
			pe := (*sender).Send(msg)
			if pe != nil {
				errs = append(errs, pe)
			}
			wg.Done()
		}(&v)
	}
	wg.Wait()
	if len(errs) > 0 {
		val := ""
		for _, v := range errs {
			val += v.Error() + "\n"
		}
		return errors.New(val)
	}
	return nil
}

func (p *senders) Init() {
	for _, sender := range p.senders {
		if pv, ok := sender.(types.Provider); ok {
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
			}
		}

		if hv, ok := sender.(types.Host); ok {
			if err := hv.Init(); err != nil {
				log.Fatalf("Host %s init failed: %s", hv.Name(), err)
				return
			}

			if err := hv.Start(); err != nil {
				log.Fatalf("Host %s start failed: %s", hv.Name(), err)
				return
			}
		}
	}
}

func (p *senders) RegisterRouter(e *gin.Engine) error {
	if len(p.senders) == 0 && !utils.IsTestInstance() {
		return errors.New("no sender found")
	}
	for _, v := range p.senders {
		if pv, ok := v.(types.SenderWithHandler); ok {
			e.Handle(pv.HandlerMethod(), pv.HandlerPath(), middleware.Auth, pv.Handler)
		}
		if pv, ok := v.(types.SenderWithSpecialHandler); ok {
			err := pv.CustomRegister(e)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
