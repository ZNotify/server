package model

import (
	"time"

	"notify-api/push/entity"
)

type Message struct {
	ID        uint   `gorm:"primary_key"`
	MessageID string `gorm:"unique"`
	UserID    string
	Title     string
	Content   string
	Long      string
	Priority  entity.Priority
	CreatedAt time.Time
}
