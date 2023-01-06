package dao

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"notify-api/ent/db"
	"notify-api/ent/generate"
	"notify-api/ent/generate/message"
	"notify-api/ent/generate/user"
	"notify-api/push/enum"
)

type messageDao struct{}

var Message = messageDao{}

var SequenceID atomic.Int64

func GetNextSequenceID() int {
	return int(SequenceID.Add(1))
}

func GetLatestMessageID(ctx context.Context) int {
	m, err := db.C.Message.Query().Order(generate.Desc(message.FieldSequenceID)).First(ctx)
	if err != nil {
		if generate.IsNotFound(err) {
			return 0
		}
		zap.S().Panicw("failed to get latest message id", "err", err)
	}
	return m.SequenceID
}

func (messageDao) CreateMessage(
	ctx context.Context,
	u *generate.User,
	id uuid.UUID,
	title string,
	content string,
	long string,
	priority enum.Priority,
	sequenceID int,
	createdAt time.Time) (*generate.Message, bool) {
	msg, err := db.C.Message.Create().
		SetUser(u).
		SetTitle(title).
		SetContent(content).
		SetLong(long).
		SetPriority(priority).
		SetSequenceID(sequenceID).
		SetID(id).
		SetCreatedAt(createdAt).
		Save(ctx)
	if err != nil {
		zap.S().Errorw("failed to create message", "err", err)
		return nil, false
	}
	return msg, true
}

// GetUserMessagesPaginated use SequenceID as afterID
func (messageDao) GetUserMessagesPaginated(ctx context.Context, u *generate.User, skip, limit int) ([]*generate.Message, bool) {
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

func (messageDao) GetUserMessagesAfterID(ctx context.Context, u *generate.User, afterID int) ([]*generate.Message, bool) {
	m, err := u.QueryMessages().
		Where(message.SequenceIDGT(afterID)).
		Order(generate.Desc(message.FieldSequenceID)).
		All(ctx)
	if err != nil {
		return nil, false
	}
	return m, true
}

func (messageDao) DeleteMessageByID(ctx context.Context, u *generate.User, messageID uuid.UUID) (int, bool) {
	row, err := db.C.Message.Delete().
		Where(message.ID(messageID)).
		Where(message.HasUserWith(user.ID(u.ID))).
		Exec(ctx)

	if err != nil {
		return 0, false
	}
	return row, true
}
