package providers

import (
	"fmt"
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
	} else {
		fmt.Println("FCMProvider registered")
	}
	err = push.Providers.Register(new(MiPushProvider))
	if err != nil {
		panic(err)
	} else {
		fmt.Println("MiPushProvider registered")
	}
	err = push.Providers.Register(new(WebPushProvider))
	if err != nil {
		panic(err)
	} else {
		fmt.Println("WebPushProvider registered")
	}
}
