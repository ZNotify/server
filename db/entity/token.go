package entity

import (
	"time"
)

type PushToken struct {
	ID        uint
	UserID    string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeviceID  string
	Channel   string
	Token     string
}
