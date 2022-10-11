package entity

import (
	"time"
)

type PushToken struct {
	ID        string
	UserID    string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeviceID  string
	Channel   string
	Token     string
}
