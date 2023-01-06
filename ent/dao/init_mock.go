//go:build test

package dao

import (
	"context"

	"notify-api/ent/db"
)

func Init() {
	db.Init()

	SequenceID.Store(int64(GetLatestMessageID(context.Background())))

	_, _ = db.C.User.Create().SetSecret("test").Save(context.Background())
}
