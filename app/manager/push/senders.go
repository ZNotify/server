package push

import (
	"github.com/pkg/errors"

	"notify-api/app/manager/push/enum"
	"notify-api/app/manager/push/host/telegram"
	"notify-api/app/manager/push/host/websocket"
	"notify-api/app/manager/push/interfaces"
	"notify-api/app/manager/push/provider/fcm"
	"notify-api/app/manager/push/provider/webpush"
	"notify-api/app/manager/push/provider/wns"
)

type senders = []interfaces.Sender

var availableSenders = senders{
	new(fcm.Provider),
	new(webpush.Provider),
	new(wns.Provider),
	new(websocket.Host),
	new(telegram.Host),
}

var activeSenders = senders{}

func IsSenderActive(id enum.Sender) bool {
	for _, v := range activeSenders {
		if v.Name() == id {
			return true
		}
	}
	return false
}

func GetSender(id enum.Sender) (interfaces.Sender, error) {
	for _, v := range availableSenders {
		if v.Name() == id {
			return v, nil
		}
	}
	return nil, errors.Errorf("sender %s not found", id)
}
