package types

import (
	"encoding/json"
	"time"

	"notify-api/db/model"

	"notify-api/push/item"
)

type Message struct {
	ID        string        `json:"id"`
	UserID    string        `json:"user_id"`
	Title     string        `json:"title"`
	Content   string        `json:"content"`
	Long      string        `json:"long"`
	Priority  item.Priority `json:"priority"`
	CreatedAt time.Time     `json:"created_at"`
}

func (m *Message) MarshalJSON() ([]byte, error) {
	type Alias Message
	return json.Marshal(&struct {
		*Alias
		CreatedAt string `json:"created_at"`
	}{
		Alias:     (*Alias)(m),
		CreatedAt: m.CreatedAt.Format(time.RFC3339),
	})
}

func FromPushMessage(msg item.PushMessage) Message {
	return Message{
		ID:        msg.MessageID,
		UserID:    msg.UserID,
		Title:     msg.Title,
		Content:   msg.Content,
		Long:      msg.Long,
		Priority:  msg.Priority,
		CreatedAt: msg.CreatedAt,
	}
}

func FromModelMessage(msg model.Message) Message {
	return Message{
		ID:        msg.MessageID,
		UserID:    msg.UserID,
		Title:     msg.Title,
		Content:   msg.Content,
		Long:      msg.Long,
		Priority:  msg.Priority,
		CreatedAt: msg.CreatedAt,
	}
}
