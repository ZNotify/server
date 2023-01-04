package item

import (
	"time"

	"notify-api/ent/generate"
)

type PushMessage struct {
	ID        string
	User      *generate.User
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
