package types

import (
	"time"

	"notify-api/serve/types"
)

type Sender interface {
	Send(msg *Message) error
	Init() error
	Name() string
}

type SenderWithConfig interface {
	Sender
	Config() any
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

// Message priority enum
const (
	PriorityLow = iota
	PriorityNormal
	PriorityHigh
)

type Message struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Long      string    `json:"long"`
	Priority  int       `json:"priority"`
	CreatedAt time.Time `json:"created_at"`
}
