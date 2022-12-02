package push

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"

	serveTypes "notify-api/serve/types"
	"notify-api/utils/config"

	"go.uber.org/zap"

	pushTypes "notify-api/push/types"
	"notify-api/utils"
)

func Send(msg *pushTypes.Message) error {
	if config.IsTest() {
		return nil
	}

	zap.S().Infof("Send message to %s", msg.UserID)

	var errs []string
	var wg sync.WaitGroup
	wg.Add(len(activeSenders))
	for _, v := range activeSenders {
		go func(sender pushTypes.Sender) {
			defer wg.Done()
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

func Init() {
	if utils.IsTestInstance() {
		activeSenders = availableSenders
		return
	}

	for id, senderCfg := range config.Config.Senders {
		sender, err := get(id)
		if err != nil {
			panic(err)
		}

		if authSender, ok := sender.(pushTypes.SenderWithConfig); ok {
			cfgType := authSender.Config()

			decodeCfg := &mapstructure.DecoderConfig{
				Metadata:         nil,
				Result:           &cfgType,
				WeaklyTypedInput: true,
				ErrorUnset:       true,
			}
			decoder, err := mapstructure.NewDecoder(decodeCfg)
			if err != nil {
				zap.S().Errorf("Init sender %s failed: %v", id, err)
			}

			err = decoder.Decode(senderCfg)
			if err != nil {
				msg := string(err.Error())
				zap.S().Fatalf("Sender %s config check failed: %s", id, msg)
			}

			config.Config.Senders[id] = cfgType
		}
		err = sender.Init()
		if err != nil {
			zap.S().Fatalf("Sender %s init failed: %v", sender.Name(), err)
		}

		if host, ok := sender.(pushTypes.Host); ok {
			err := host.Start()
			if err != nil {
				zap.S().Fatalf("Sender %s start failed: %v", sender.Name(), err)
			}
		}

		activeSenders = append(activeSenders, sender)

	}
}

func RegisterRouter(e *gin.RouterGroup) {
	for _, v := range activeSenders {
		if pv, ok := v.(pushTypes.SenderWithHandler); ok {
			e.Handle(pv.HandlerMethod(), pv.HandlerPath(), serveTypes.WrapHandler(pv.Handler))
		}
	}
}
