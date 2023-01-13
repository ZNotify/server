package entity

import (
	"notify-api/push/enum"
)

type Device struct {
	Identifier string      `json:"identifier"`
	Channel    enum.Sender `json:"channel"`
	DeviceName string      `json:"deviceName"`
	DeviceMeta string      `json:"deviceMeta"`
}
