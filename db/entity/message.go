package entity

import (
	"time"

	"notify-api/push/types"
)

type Message struct {
	ID        uint   `gorm:"primary_key"`
	MessageID string `gorm:"unique"`
	UserID    string
	Title     string
	Content   string
	Long      string
	Priority  types.Priority
	CreatedAt time.Time
}
