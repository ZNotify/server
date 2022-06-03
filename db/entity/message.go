package entity

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

type Message struct {
	ID        string
	UserID    string
	Title     string
	Content   string
	Long      string
	CreatedAt time.Time
}

func (m Message) ToJSON() (string, error) {
	data := m.ToGinH()
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (m Message) ToGinH() gin.H {
	return gin.H{
		"id":         m.ID,
		"user_id":    m.UserID,
		"title":      m.Title,
		"content":    m.Content,
		"long":       m.Long,
		"created_at": m.CreatedAt.Format(time.RFC3339),
	}
}
