package entity

import "time"

type Device struct {
	ID        uint   `gorm:"primary_key"`
	DeviceID  string `gorm:"unique"`
	Meta      string
	Channel   string
	Token     string
	CreatedAt time.Time
	UserID    string
}
