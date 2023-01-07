package dao

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"notify-api/ent/db"
	"notify-api/ent/generate"
	"notify-api/ent/generate/device"
	"notify-api/ent/generate/user"
)

type userDao struct{}

var User = userDao{}

func (userDao) EnsureUser(ctx context.Context, githubID int64, githubName, githubLogin, githubOAuthToken string) (*generate.User, bool) {
	u, err := db.C.User.Query().Where(user.GithubID(githubID)).Only(ctx)
	if err != nil {
		if !generate.IsNotFound(err) {
			zap.S().Errorw("failed to check user exists", "err", err)
			return nil, false
		}
		// Create user
		u, err = db.C.User.Create().
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
		_, err = db.C.User.UpdateOne(u).
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
	d, err := db.C.Device.
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
	u, err := db.C.User.Query().Where(user.Secret(secret)).Only(ctx)
	if err != nil {
		if !generate.IsNotFound(err) {
			zap.S().Errorw("failed to check secret valid", "err", err)
		}
		return nil, false
	}
	return u, true
}
