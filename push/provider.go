package push

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Provider interface {
	Send(msg *Message) error
	Check() error
	Init(e *gin.Engine) error
}

type Message struct {
	ID        string
	UserID    string
	Title     string
	Content   string
	Long      string
	CreatedAt time.Time
}
