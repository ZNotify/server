package push

import (
	"errors"
	"github.com/gin-gonic/gin"
	"notify-api/push/providers"
	"notify-api/utils"
	"sync"
	"time"
)

type Provider interface {
	Send(msg *Message) error
	Check() error
	Init(e *gin.Engine) error
}

type Message struct {
	ID        string
	UserID    string
	Title     string
	Content   string
	Long      string
	CreatedAt time.Time
}

var ps = []Provider{new(providers.FCMProvider), new(providers.MiPushProvider), new(providers.WebPushProvider)}

func Init(e *gin.Engine) {
	if utils.IsTestInstance() {
		return
	}
	for _, v := range ps {
		err := v.Check()
		if err != nil {
			panic(err)
		}
	}
	wg := sync.WaitGroup{}
	wg.Add(len(ps))
	for _, v := range ps {
		p := &v
		go func() {
			err := (*p).Init(e)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func Send(msg *Message) error {
	var errs []error
	var wg sync.WaitGroup
	wg.Add(len(ps))
	for _, v := range ps {
		provider := &v
		go func() {
			pe := (*provider).Send(msg)
			if pe != nil {
				errs = append(errs, pe)
			}
			wg.Done()
		}()
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
