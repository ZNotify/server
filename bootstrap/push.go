package bootstrap

import "notify-api/app/manager/push"

func initializePushManager() {
	push.Init()
}
