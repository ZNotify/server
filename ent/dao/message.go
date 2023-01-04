package dao

import (
	"context"

	"github.com/google/uuid"

	"notify-api/ent/db"
	"notify-api/ent/generate"
	"notify-api/ent/generate/message"
	"notify-api/ent/generate/user"
)

type messageDao struct{}

var Message = messageDao{}

// GetUserMessagesPaginated use sequenceID as afterID
func (messageDao) GetUserMessagesPaginated(ctx context.Context, u *generate.User, skip int, limit int) ([]*generate.Message, bool) {
	messages, err := u.QueryMessages().
		Order(generate.Desc(message.FieldSequenceID)).
		Offset(skip).
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, false
	}
	return messages, true
}

func (messageDao) GetUserMessage(ctx context.Context, u *generate.User, messageID uuid.UUID) (*generate.Message, bool) {
	m, err := u.QueryMessages().
		Where(message.ID(messageID)).
		Only(ctx)
	if err != nil {
		if generate.IsNotFound(err) {
			return nil, true
		}
		return nil, false
	}
	return m, true
}

func (messageDao) DeleteMessageByID(ctx context.Context, u *generate.User, messageID uuid.UUID) (int, bool) {
	row, err := db.Client.Message.Delete().
		Where(message.ID(messageID)).
		Where(message.HasUserWith(user.ID(u.ID))).
		Exec(ctx)

	if err != nil {
		return 0, false
	}
	return row, true
}
