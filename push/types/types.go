package types

import (
	"time"

	"notify-api/serve/types"
)

type Config = map[string]string

type Sender interface {
	Send(msg *Message) error
	Name() string
}

type SenderWithoutConfig interface {
	Sender
	Init() error
}

type SenderWithConfig interface {
	Sender
	Init(config Config) error
	Config() []string
}

type Host interface {
	Sender
	Start() error
}

type SenderWithHandler interface {
	Sender
	Handler(ctx *types.Ctx)
	HandlerPath() string
	HandlerMethod() string
}

type SenderWithInitialTokenMeta interface {
	Sender
	GetInitialTokenMeta() string
}

type Priority string

const (
	PriorityLow    Priority = "low"    // low
	PriorityNormal Priority = "normal" // normal
	PriorityHigh   Priority = "high"   // high
)

type Message struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Long      string    `json:"long"`
	Priority  Priority  `json:"priority"`
	CreatedAt time.Time `json:"created_at"`
}
