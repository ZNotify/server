package safeMessage

import (
	"notify-api/db/entity"
	"time"
)

type SafeMessage struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Long      string `json:"long"`
	CreatedAt string `json:"created_at"`
}

func FromEntityMessage(msg entity.Message) SafeMessage {
	return SafeMessage{
		ID:        msg.ID,
		UserID:    msg.UserID,
		Title:     msg.Title,
		Content:   msg.Content,
		Long:      msg.Long,
		CreatedAt: msg.CreatedAt.Format(time.RFC3339),
	}
}

func FromEntityMessageArray(msgs []entity.Message) []SafeMessage {
	var safeMsgs []SafeMessage
	for _, msg := range msgs {
		safeMsgs = append(safeMsgs, FromEntityMessage(msg))
	}
	return safeMsgs
}
