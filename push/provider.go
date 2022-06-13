package push

import (
	"errors"
	"fmt"
	"github.com/ZNotify/server/db/entity"
	"github.com/ZNotify/server/utils"
	"github.com/gin-gonic/gin"
	"os"
	"sync"
)

type Provider interface {
	send(msg *entity.Message) error
	check() error
	init(e *gin.Engine) (Provider, error)
}

var providers = []Provider{new(FCMProvider), new(WebPushProvider), new(MiPushProvider)}

func Init(e *gin.Engine) {
	if utils.IsTestInstance() {
		return
	}
	for _, v := range providers {
		_, err := v.init(e)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func Send(msg *entity.Message) error {
	var errs []error
	var wg sync.WaitGroup
	wg.Add(len(providers))
	for _, v := range providers {
		provider := v
		go func() {
			pe := provider.send(msg)
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
