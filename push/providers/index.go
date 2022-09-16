package providers

import (
	"fmt"
	"notify-api/push"
)

func Init() {
	var err error

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
