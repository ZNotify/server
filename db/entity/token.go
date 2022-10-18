package entity

import (
	"time"
)

type PushToken struct {
	ID        uint `gorm:"primary_key"`
	UserID    string
	CreatedAt time.Time
	DeviceID  string `gorm:"unique_index:idx_device_id"`
	Channel   string
	Token     string
}
