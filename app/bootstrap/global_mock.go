//go:build test

package bootstrap

import (
	"context"

	"go.uber.org/zap"

	"notify-api/app/db/dao"
	"notify-api/app/global"
)

func initializeGlobalVar() {
	dao.SequenceID.Store(int64(dao.GetLatestMessageID(context.Background())))

	u, err := global.App.DB.User.Create().
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
