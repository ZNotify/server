package providers

import (
	"fmt"
	"notify-api/push"
)

func Init() {
	var err error

	err = push.Providers.Register(new(HostPushProvider))
	if err != nil {
		panic(err)
	} else {
		fmt.Println("HostPushProvider registered")
	}

	err = push.Providers.Register(new(FCMProvider))
	if err != nil {
		panic(err)
	} else {
		fmt.Println("FCMProvider registered")
	}

	err = push.Providers.Register(new(WebPushProvider))
	if err != nil {
		panic(err)
	} else {
		fmt.Println("WebPushProvider registered")
	}
}
