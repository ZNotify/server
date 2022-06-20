package providers

import (
	"notify-api/push"
	"notify-api/utils"
)

func Init() {
	if utils.IsTestInstance() {
		return
	}

	var err error

	err = push.Providers.Register(new(FCMProvider))
	if err != nil {
		panic(err)
	}
	err = push.Providers.Register(new(MiPushProvider))
	if err != nil {
		panic(err)
	}
	err = push.Providers.Register(new(WebPushProvider))
	if err != nil {
		panic(err)
	}
}
