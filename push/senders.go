package push

import (
	"notify-api/push/host"
	"notify-api/push/provider"
	pushTypes "notify-api/push/types"
)

type senders = []pushTypes.Sender

var availableSenders = senders{
	new(provider.FCMProvider),
	new(provider.WebPushProvider),
	new(host.WebSocketHost),
	new(host.TelegramHost),
}

var activeSenders = senders{}

func IsValid(id string) bool {
	for _, v := range availableSenders {
		if v.Name() == id {
			return true
		}
	}
	return false
}
