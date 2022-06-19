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
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Long      string    `json:"long"`
	CreatedAt time.Time `json:"created_at"`
}
