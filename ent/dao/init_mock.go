//go:build test

package dao

import (
	"context"

	"go.uber.org/zap"

	"notify-api/ent/db"
)

func Init() {
	db.Init()

	SequenceID.Store(int64(GetLatestMessageID(context.Background())))

	u, err := db.C.User.Create().
		SetSecret("test").
		SetGithubID(1).
		SetGithubName("test").
		SetGithubLogin("test").
		SetGithubOauthToken("test").
		Save(context.Background())

	if err != nil {
		panic(err)
	} else {
		zap.S().Infof("Created test user: %v", u)
	}
}
