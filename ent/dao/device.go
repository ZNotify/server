package dao

import (
	"context"

	"golang.org/x/exp/slices"

	"go.uber.org/zap"

	"notify-api/ent/db"
	"notify-api/ent/generate"
	"notify-api/ent/generate/device"
)

type deviceDao struct{}

var DeviceDao = deviceDao{}

func (deviceDao) EnsureDevice(
	ctx context.Context,
	identifier string,
	channel string,
	channelMeta string,
	channelToken string,
	deviceName string,
	deviceMeta string,
	u *generate.User) (*generate.Device, bool) {

	id, err := db.Client.Device.
		Create().
		SetIdentifier(identifier).
		SetChannel(channel).
		SetChannelMeta(channelMeta).
		SetChannelToken(channelToken).
		SetDeviceName(deviceName).
		SetDeviceMeta(deviceMeta).
		SetUser(u).
		OnConflictColumns("id", "identifier").
		UpdateNewValues().
		Update(func(upsert *generate.DeviceUpsert) {
			uc := upsert.UpdateSet.UpdateColumns()
			if slices.Contains(uc, "channel") {
				upsert.SetChannelMeta("")
			}
		}).
		ID(ctx)

	if err != nil {
		zap.S().Errorw("failed to create device", "err", err)
		return nil, false
	}

	return db.Client.Device.GetX(ctx, id), true
}

func (deviceDao) GetDeviceByIdentifier(ctx context.Context, identifier string) (*generate.Device, bool) {
	d, err := db.Client.Device.
		Query().
		Where(device.Identifier(identifier)).
		Only(ctx)
	if err != nil {
		if !generate.IsNotFound(err) {
			zap.S().Errorw("failed to get device by identifier", "err")
		}
		return nil, false
	}

	return d, true
}

func (userDao) GetUserDeviceByIdentifier(ctx context.Context, u *generate.User, identifier string) (*generate.Device, bool) {
	d, err := u.QueryDevices().Where(device.Identifier(identifier)).Only(ctx)
	if err != nil {
		if !generate.IsNotFound(err) {
			zap.S().Errorw("failed to get user device by identifier", "err", err)
		}
		return nil, false
	}
	return d, true
}

func (deviceDao) DeleteDeviceByIdentifier(ctx context.Context, identifier string) bool {
	_, err := db.Client.Device.Delete().Where(device.Identifier(identifier)).Exec(ctx)
	if err != nil {
		if !generate.IsNotFound(err) {
			zap.S().Errorw("failed to delete device by identifier", "err", err)
		}
		return false
	}
	return true
}

func (deviceDao) DeleteDevice(ctx context.Context, d *generate.Device) bool {
	err := db.Client.Device.DeleteOne(d).Exec(ctx)
	if err != nil {
		if !generate.IsNotFound(err) {
			zap.S().Errorw("failed to delete device", "err", err)
		}
		return false
	}
	return true
}
