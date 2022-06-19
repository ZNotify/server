package providers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"notify-api/push"
	"notify-api/utils"
	"sync"
)

var ps = []push.Provider{new(FCMProvider), new(MiPushProvider), new(WebPushProvider)}

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
		go func(p push.Provider) {
			err := p.Init(e)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(v)
	}
	wg.Wait()
}

func Send(msg *push.Message) error {
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
