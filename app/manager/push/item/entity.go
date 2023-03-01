package item

import (
	"time"

	"notify-api/app/db/dao"
	"notify-api/app/db/ent/generate"
	"notify-api/app/manager/push/enum"

	"github.com/google/uuid"
)

type PushMessage struct {
	ID         uuid.UUID
	User       *generate.User
	Title      string
	Content    string
	Long       string
	Priority   enum.Priority
	CreatedAt  time.Time
	SequenceID int
}

func NewPushMessage(user *generate.User, title, content, long string, priority enum.Priority) *PushMessage {
	return &PushMessage{
		ID:         uuid.New(),
		User:       user,
		Title:      title,
		Content:    content,
		Long:       long,
		Priority:   priority,
		CreatedAt:  time.Now(),
		SequenceID: dao.GetNextSequenceID(),
	}
}

func FromModelMessageWithUser(msg *generate.Message, u *generate.User) *PushMessage {
	return &PushMessage{
		ID:         msg.ID,
		User:       u,
		Title:      msg.Title,
		Content:    msg.Content,
		Long:       msg.Long,
		Priority:   msg.Priority,
		CreatedAt:  msg.CreatedAt,
		SequenceID: msg.SequenceID,
	}
}
