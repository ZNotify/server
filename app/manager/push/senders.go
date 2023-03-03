package push

import (
	"github.com/pkg/errors"

	"github.com/ZNotify/server/app/manager/push/enum"
	"github.com/ZNotify/server/app/manager/push/host/telegram"
	"github.com/ZNotify/server/app/manager/push/host/websocket"
	"github.com/ZNotify/server/app/manager/push/interfaces"
	"github.com/ZNotify/server/app/manager/push/provider/fcm"
	"github.com/ZNotify/server/app/manager/push/provider/webpush"
	"github.com/ZNotify/server/app/manager/push/provider/wns"
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
