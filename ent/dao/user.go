package dao

import (
	"context"

	"go.uber.org/zap"

	"notify-api/ent/db"
	"notify-api/ent/generate"
	"notify-api/ent/generate/device"
	"notify-api/ent/generate/user"
)

type userDao struct{}

var User = userDao{}

func (userDao) GetDeviceUser(ctx context.Context, identifier string) (*generate.User, bool) {
	d, err := db.Client.Device.
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
	u, err := db.Client.User.Query().Where(user.Secret(secret)).Only(ctx)
	if err != nil {
		if !generate.IsNotFound(err) {
			zap.S().Errorw("failed to check secret valid", "err", err)
		}
		return nil, false
	}
	return u, true
}
