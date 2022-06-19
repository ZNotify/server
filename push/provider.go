package push

import (
	"errors"
	"github.com/gin-gonic/gin"
	"notify-api/db/entity"
	"notify-api/utils"
	"sync"
)

type Provider interface {
	send(msg *entity.Message) error
	check() error
	init(e *gin.Engine) error
}

var providers = []Provider{new(FCMProvider), new(WebPushProvider), new(MiPushProvider)}

func Init(e *gin.Engine) {
	if utils.IsTestInstance() {
		return
	}
	for _, v := range providers {
		err := v.check()
		if err != nil {
			panic(err)
		}
	}
	wg := sync.WaitGroup{}
	wg.Add(len(providers))
	for _, v := range providers {
		providers := &v
		go func() {
			err := (*providers).init(e)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func Send(msg *entity.Message) error {
	var errs []error
	var wg sync.WaitGroup
	wg.Add(len(providers))
	for _, v := range providers {
		provider := &v
		go func() {
			pe := (*provider).send(msg)
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
