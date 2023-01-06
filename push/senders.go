package push

import (
	"github.com/pkg/errors"

	"notify-api/push/host/telegram"
	"notify-api/push/host/websocket"
	"notify-api/push/provider/fcm"
	"notify-api/push/provider/webpush"
	"notify-api/push/provider/wns"
	pushTypes "notify-api/push/types"
)

type senders = []pushTypes.Sender

var availableSenders = senders{
	new(fcm.Provider),
	new(webpush.Provider),
	new(wns.Provider),
	new(websocket.Host),
	new(telegram.Host),
}

var activeSenders = senders{}

func IsSenderIdValid(id string) bool {
	for _, v := range activeSenders {
		if v.Name() == id {
			return true
		}
	}
	return false
}

func GetSender(id string) (pushTypes.Sender, error) {
	for _, v := range availableSenders {
		if v.Name() == id {
			return v, nil
		}
	}
	return nil, errors.Errorf("sender %s not found", id)
}
