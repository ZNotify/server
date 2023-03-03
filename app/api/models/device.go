package models

import (
	"github.com/ZNotify/server/app/manager/push/enum"
)

type Device struct {
	Identifier string      `json:"identifier"`
	Channel    enum.Sender `json:"channel"`
	DeviceName string      `json:"deviceName"`
	DeviceMeta string      `json:"deviceMeta"`
}
