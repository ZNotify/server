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

type SenderAuth = map[string]string

type SenderWithAuth interface {
	Sender
	Check(SenderAuth) error
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

type Message struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Long      string    `json:"long"`
	CreatedAt time.Time `json:"created_at"`
}
