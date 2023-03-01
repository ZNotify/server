//go:build !test

package push

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"notify-api/app/db/helper"
	"notify-api/app/global"
	"notify-api/app/manager/push/enum"
	pushTypes "notify-api/app/manager/push/interfaces"
	"notify-api/app/manager/push/item"
	"notify-api/app/utils"

	"go.uber.org/zap"
)

func Send(ctx context.Context, msg *item.PushMessage) error {
	zap.S().Infof("Send message to %s", helper.GetReadableName(msg.User))

	var errs []string
	var wg sync.WaitGroup
	wg.Add(len(activeSenders))
	for _, v := range activeSenders {
		go func(sender pushTypes.Sender) {
			defer wg.Done()
			pe := sender.Send(ctx, msg)
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
	for id, senderCfg := range global.App.Config.Senders {
		if IsSenderActive(enum.Sender(id)) {
			zap.S().Fatalf("Sender %s load twice", id)
		}

		sender, err := GetSender(enum.Sender(id))
		if err != nil {
			zap.S().Fatalf("Failed to get sender %s: %v", id, err)
		}

		if authSender, ok := sender.(pushTypes.SenderWithConfig); ok {
			cfgKeys := authSender.Config()

			cfg := make(map[string]string)
			for _, v := range cfgKeys {
				value, ok := senderCfg[v]
				if !ok {
					zap.S().Fatalf("Sender %s config %s not found", id, v)
				}
				cfg[v] = utils.YamlStringClean(value)
			}
			err = authSender.Init(cfg)
			if err != nil {
				zap.S().Fatalf("Sender %s init failed: %v", id, err)
			}
		} else {
			cs, ok := sender.(pushTypes.SenderWithoutConfig)
			if ok {
				err = cs.Init()
				if err != nil {
					zap.S().Fatalf("Sender %s init failed: %v", id, err)
				}
			}
		}

		if host, ok := sender.(pushTypes.SenderWithBackground); ok {
			err := host.Setup()
			if err != nil {
				zap.S().Fatalf("Sender %s start failed: %v", sender.Name(), err)
			}
		}

		activeSenders = append(activeSenders, sender)
	}
}
