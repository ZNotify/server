package item

import (
	"time"
)

type PushMessage struct {
	MessageID string
	UserID    string
	Title     string
	Content   string
	Long      string
	Priority  Priority
	CreatedAt time.Time
}

type Priority string

const (
	PriorityLow    Priority = "low"    // low
	PriorityNormal Priority = "normal" // normal
	PriorityHigh   Priority = "high"   // high
)
