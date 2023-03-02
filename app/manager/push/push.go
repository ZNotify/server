//go:build !test

package push

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"

	"notify-api/app/db/helper"
	"notify-api/app/global"
	"notify-api/app/manager/push/enum"
	"notify-api/app/manager/push/interfaces"
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
		go func(sender interfaces.Sender) {
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
	cfgSendersV := reflect.ValueOf(global.App.Config.Senders)
	if cfgSendersV.IsZero() {
		zap.S().Fatalf("No sender found in config")
	}

	for k := 0; k < cfgSendersV.NumField(); k++ {
		var senderCfg any
		field := cfgSendersV.Type().Field(k)
		senderName := field.Name
		senderCfgField := cfgSendersV.Field(k)
		senderCfg = senderCfgField.Interface()

		if IsSenderActive(enum.Sender(senderName)) {
			zap.S().Fatalf("Sender %s load twice", senderName)
		}

		sender, err := GetSender(enum.Sender(senderName))
		if err != nil {
			zap.S().Fatalf("Failed to get sender %s: %v", senderName, err)
		}

		if cfgSender, ok := sender.(interfaces.SenderWithConfig); ok {
			if senderCfgField.IsZero() {
				zap.S().Infof("Sender %s is disabled", senderName)
				continue
			}

			err = cfgSender.Init(senderCfg)
			if err != nil {
				zap.S().Fatalf("Sender %s init failed: %v", senderName, err)
			}
		} else {
			cs, ok := sender.(interfaces.SenderWithoutConfig)
			if ok {
				enable := senderCfg.(bool)
				if !enable {
					zap.S().Infof("Sender %s is disabled", senderName)
					continue
				}

				err = cs.Init()
				if err != nil {
					zap.S().Fatalf("Sender %s init failed: %v", senderName, err)
				}
			}
		}

		if host, ok := sender.(interfaces.SenderWithBackground); ok {
			err = host.Setup()
			if err != nil {
				zap.S().Fatalf("Sender %s start failed: %v", sender.Name(), err)
			}
		}

		activeSenders = append(activeSenders, sender)
		zap.S().Infof("Sender %s is loaded", senderName)
	}

	if len(activeSenders) == 0 {
		zap.S().Fatalf("No sender enabled")
	}
}
