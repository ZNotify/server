package dao

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"notify-api/db/ent/generate"
	"notify-api/db/ent/generate/device"
	"notify-api/db/ent/generate/user"
	"notify-api/global"
)

type userDao struct{}

var User = userDao{}

func (userDao) EnsureUser(ctx context.Context, githubID int64, githubName, githubLogin, githubOAuthToken string) (*generate.User, bool) {
	u, err := global.App.DB.User.Query().Where(user.GithubID(githubID)).Only(ctx)
	if err != nil {
		if !generate.IsNotFound(err) {
			zap.S().Errorw("failed to check user exists", "err", err)
			return nil, false
		}
		// Create user
		u, err = global.App.DB.User.Create().
			SetSecret(uuid.New().String()).
			SetGithubID(githubID).
			SetGithubName(githubName).
			SetGithubLogin(githubLogin).
			SetGithubOauthToken(githubOAuthToken).
			Save(ctx)
		if err != nil {
			zap.S().Errorw("failed to create user", "err", err)
			return nil, false
		}
		return u, true
	} else {
		// Update user
		_, err = global.App.DB.User.UpdateOne(u).
			SetGithubName(githubName).
			SetGithubLogin(githubLogin).
			SetGithubOauthToken(githubOAuthToken).
			Save(ctx)
		if err != nil {
			zap.S().Errorw("failed to update user", "err", err)
			return nil, false
		}
		return u, true
	}
}

func (userDao) GetDeviceUser(ctx context.Context, identifier string) (*generate.User, bool) {
	d, err := global.App.DB.Device.
		Query().
		Where(device.Identifier(identifier)).
		WithUser().
		Only(ctx)

	if err != nil {
		if !generate.IsNotFound(err) {
			zap.S().Errorw("failed to get device user", "err")
		}
		return nil, false
	}

	return d.Edges.User, true
}

func (userDao) GetUserBySecret(ctx context.Context, secret string) (*generate.User, bool) {
	u, err := global.App.DB.User.Query().Where(user.Secret(secret)).Only(ctx)
	if err != nil {
		if !generate.IsNotFound(err) {
			zap.S().Errorw("failed to check secret valid", "err", err)
		}
		return nil, false
	}
	return u, true
}
