package bootstrap

import "github.com/ZNotify/server/app/manager/push"

func initializePushManager() {
	push.Init()
}
