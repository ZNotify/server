//go:build test

package push

import pushTypes "notify-api/push/types"

func Send(msg *pushTypes.Message) error {
	return nil
}

func Init() {
	activeSenders = availableSenders
}
