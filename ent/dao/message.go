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
	"notify-api/push/item"
)

type messageDao struct{}

var Message = messageDao{}

var sequenceID atomic.Int64

func init() {
	sequenceID.Store(int64(GetLatestMessageID(context.Background())))
}

func GetLatestMessageID(ctx context.Context) int {
	m, err := db.Client.Message.Query().Order(generate.Desc(message.FieldSequenceID)).First(ctx)
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
	priority item.Priority,
	createdAt time.Time) (*generate.Message, bool) {
	seq := sequenceID.Add(1)
	msg, err := db.Client.Message.Create().
		SetUser(u).
		SetTitle(title).
		SetContent(content).
		SetLong(long).
		SetPriority(priority).
		SetSequenceID(int(seq)).
		SetID(id).
		SetCreatedAt(createdAt).
		Save(ctx)
	if err != nil {
		zap.S().Errorw("failed to create message", "err", err)
		return nil, false
	}
	return msg, true
}

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
