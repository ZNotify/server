package dao

import (
	"context"

	"go.uber.org/zap"

	"notify-api/ent/db"
	"notify-api/ent/generate"
	"notify-api/ent/generate/device"
	"notify-api/push/enum"
)

type deviceDao struct{}

var Device = deviceDao{}

// EnsureDevice ensures that a device exists for the given user and identifier.
func (dd deviceDao) EnsureDevice(
	ctx context.Context,
	identifier string,
	channel enum.Sender,
	channelToken string,
	deviceName string,
	deviceMeta string,
	u *generate.User) (d *generate.Device, isChannelChange bool, oldDevice *generate.Device, ok bool) {
	oldDevice, ok = dd.GetUserDeviceByIdentifier(ctx, u, identifier)
	if ok {
		// Update device
		_, err := db.C.Device.UpdateOne(oldDevice).
			SetIdentifier(identifier).
			SetChannel(channel).
			SetChannelToken(channelToken).
			SetDeviceName(deviceName).
			SetDeviceMeta(deviceMeta).
			SetUser(u).
			Save(ctx)
		if err != nil {
			zap.S().Errorw("failed to update device", "err", err)
			return nil, false, oldDevice, false
		}
		if oldDevice.Channel != channel {
			return oldDevice, true, oldDevice, true
		} else {
			return oldDevice, false, nil, true
		}
	} else {
		// Create device
		d, err := db.C.Device.Create().
			SetIdentifier(identifier).
			SetChannel(channel).
			SetChannelToken(channelToken).
			SetDeviceName(deviceName).
			SetDeviceMeta(deviceMeta).
			SetUser(u).
			Save(ctx)
		if err != nil {
			zap.S().Errorw("failed to create device", "err", err)
			return nil, false, nil, false
		}
		return d, false, nil, true
	}
}

func (deviceDao) UpdateDeviceChannelMeta(ctx context.Context, d *generate.Device, channelMeta string) bool {
	_, err := db.C.Device.UpdateOne(d).SetChannelMeta(channelMeta).Save(ctx)
	if err != nil {
		zap.S().Errorw("failed to update device channel meta", "err", err)
		return false
	}
	return true
}

func (deviceDao) GetDeviceByIdentifier(ctx context.Context, identifier string) (*generate.Device, bool) {
	d, err := db.C.Device.
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

func (deviceDao) GetUserDevices(ctx context.Context, u *generate.User) ([]*generate.Device, bool) {
	ds, err := u.QueryDevices().All(ctx)
	if err != nil {
		zap.S().Errorw("failed to get user devices", "err", err)
		return nil, false
	}
	return ds, true
}

func (deviceDao) GetUserDeviceByIdentifier(ctx context.Context, u *generate.User, identifier string) (*generate.Device, bool) {
	d, err := u.QueryDevices().Where(device.Identifier(identifier)).WithUser().Only(ctx)
	if err != nil {
		if !generate.IsNotFound(err) {
			zap.S().Errorw("failed to get user device by identifier", "err", err)
		}
		return nil, false
	}
	return d, true
}

func (deviceDao) GetUserDeviceChannelTokens(ctx context.Context, u *generate.User, channel enum.Sender) ([]string, bool) {
	devices := make([]string, 0)
	err := u.
		QueryDevices().
		Where(device.Channel(channel)).
		Select(device.FieldChannelToken).
		Scan(ctx, &devices)
	if err != nil {
		zap.S().Errorw("failed to get user device channel tokens", "err", err)
		return nil, false
	}
	return devices, true
}

func (deviceDao) DeleteDeviceByIdentifier(ctx context.Context, identifier string) bool {
	_, err := db.C.Device.Delete().Where(device.Identifier(identifier)).Exec(ctx)
	if err != nil {
		if !generate.IsNotFound(err) {
			zap.S().Errorw("failed to delete device by identifier", "err", err)
		}
		return false
	}
	return true
}

func (deviceDao) DeleteDevice(ctx context.Context, d *generate.Device) bool {
	err := db.C.Device.DeleteOne(d).Exec(ctx)
	if err != nil {
		if !generate.IsNotFound(err) {
			zap.S().Errorw("failed to delete device", "err", err)
		}
		return false
	}
	return true
}
