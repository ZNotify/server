package push

import (
	"notify-api/push/host/telegram"
	"notify-api/push/host/websocket"
	"notify-api/push/provider/fcm"
	"notify-api/push/provider/webpush"
	pushTypes "notify-api/push/types"
)

type senders = []pushTypes.Sender

var availableSenders = senders{
	new(fcm.Provider),
	new(webpush.Provider),
	new(websocket.Host),
	new(telegram.Host),
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
