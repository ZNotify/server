package entity

import (
	"time"
)

type PushToken struct {
	ID        uint `gorm:"primary_key"`
	UserID    string
	CreatedAt time.Time
	DeviceID  string `gorm:"unique"`
	Channel   string
	Token     string
	TokenMeta string
}
