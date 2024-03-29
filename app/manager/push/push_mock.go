//go:build test

package push

import (
	"context"

	"go.uber.org/zap"

	"github.com/ZNotify/server/app/db/helper"
	"github.com/ZNotify/server/app/manager/push/item"
)

func Send(ctx context.Context, msg *item.PushMessage) error {
	fields := []zap.Field{
		zap.String("user_id", helper.GetReadableName(msg.User)),
		zap.String("title", msg.Title),
		zap.String("content", msg.Content),
		zap.String("long_content", msg.Long),
		zap.String("priority", string(msg.Priority)),
		zap.Int("sequence_id", msg.SequenceID),
	}

	zap.L().Info("Try to send message", fields...)
	return nil
}

func Init() {
	activeSenders = availableSenders
}
