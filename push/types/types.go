package types

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Sender interface {
	Send(msg *Message) error
	Init() error
	Name() string
}

type Provider interface {
	Sender
	Check() error
}

type Host interface {
	Sender
	Start() error
}

type SenderWithHandler interface {
	Sender
	Handler(ctx *gin.Context)
	HandlerPath() string
	HandlerMethod() string
}

type SenderWithSpecialHandler interface {
	CustomRegister(ctx *gin.Engine) error
}

type Message struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Long      string    `json:"long"`
	CreatedAt time.Time `json:"created_at"`
}
