package push

import (
	"errors"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

type Provider interface {
	Send(msg *Message) error
	Check() error
	Init() error
	ChannelName() string
}

type ProvidersWithHandler interface {
	ProviderHandler(ctx *gin.Context)
	ProviderHandlerPath() string
	ProviderHandlerMethod() string
}

type Message struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Long      string    `json:"long"`
	CreatedAt time.Time `json:"created_at"`
}

type providers struct {
	providerMap map[string]Provider
}

var Providers = providers{
	providerMap: make(map[string]Provider),
}

func (p *providers) Send(msg *Message) error {
	var errs []error
	var wg sync.WaitGroup
	wg.Add(len(p.providerMap))
	for _, v := range p.providerMap {
		go func(provider Provider) {
			pe := provider.Send(msg)
			if pe != nil {
				errs = append(errs, pe)
			}
			wg.Done()
		}(v)
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

func (p *providers) Register(pv Provider) error {
	name := pv.ChannelName()
	if _, ok := p.providerMap[name]; ok {
		return errors.New("providerMap already registered")
	}

	err := pv.Check()
	if err != nil {
		return err
	}
	err = pv.Init()
	if err != nil {
		return err
	}

	p.providerMap[name] = pv
	return nil
}

func (p *providers) RegisterRouter(e *gin.Engine) error {
	if len(p.providerMap) == 0 {
		return errors.New("providerMap is empty")
	}
	for _, v := range p.providerMap {
		if pv, ok := v.(ProvidersWithHandler); ok {
			e.Handle(pv.ProviderHandlerMethod(), pv.ProviderHandlerPath(), pv.ProviderHandler)
		}
	}
	return nil
}
