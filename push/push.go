//go:build !test

package push

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"notify-api/push/item"
	"notify-api/utils/config"

	"go.uber.org/zap"

	pushTypes "notify-api/push/types"
	"notify-api/utils"
)

func Send(msg *item.PushMessage) error {
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
	for id, senderCfg := range config.Config.Senders {
		sender, err := get(id)
		if err != nil {
			panic(err)
		}

		if authSender, ok := sender.(pushTypes.SenderWithConfig); ok {
			cfgKeys := authSender.Config()

			cfg := make(map[string]string)
			for _, v := range cfgKeys {
				value, ok := senderCfg[v]
				if !ok {
					zap.S().Fatalf("Sender %s config %s not found", id, v)
				}
				cfg[v] = utils.TokenClean(value)
			}
			err = authSender.Init(cfg)
			if err != nil {
				zap.S().Fatalf("Sender %s init failed: %v", id, err)
			}
		} else {
			err := sender.(pushTypes.SenderWithoutConfig).Init()
			if err != nil {
				zap.S().Fatalf("Sender %s init failed: %v", id, err)
			}
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
