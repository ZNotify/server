package util

import (
	. "notify-api/db"
	"notify-api/db/model"
)

type deviceUtil struct{}

var DeviceUtil = deviceUtil{}

func (deviceUtil) CreateOrUpdate(userID string, deviceID string, channel string, token string, info string) error {
	var d model.Device
	RWLock.RLock()
	ret := DB.
		Where(model.Device{DeviceID: deviceID}).
		FirstOrInit(&d)
	RWLock.RUnlock()
	if ret.Error != nil {
		return ret.Error
	}

	d.Meta = ""
	d.Channel = channel
	d.Token = token
	d.UserID = userID
	d.DeviceInfo = info

	RWLock.Lock()
	ret = DB.Save(&d)
	RWLock.Unlock()

	return ret.Error
}

func (deviceUtil) GetDeviceMeta(deviceID string) (string, error) {
	var d model.Device
	RWLock.RLock()
	ret := DB.
		Where(model.Device{DeviceID: deviceID}).
		Select("meta").
		First(&d)
	RWLock.RUnlock()
	if ret.Error != nil {
		return "", ret.Error
	}
	return d.Meta, nil
}

func (deviceUtil) UpdateDeviceMeta(deviceID string, meta string) error {
	RWLock.Lock()
	ret := DB.
		Model(&model.Device{}).
		Where(model.Device{DeviceID: deviceID}).
		Update("meta", meta)
	RWLock.Unlock()
	return ret.Error
}

func (deviceUtil) GetUserChannelTokens(userID string, channel string) ([]string, error) {
	var tokens []string
	RWLock.RLock()
	ret := DB.
		Model(&model.Device{}).
		Where(model.Device{UserID: userID, Channel: channel}).
		Pluck("token", &tokens)
	RWLock.RUnlock()
	if ret.Error != nil {
		return nil, ret.Error
	}
	return tokens, nil
}

func (deviceUtil) GetDeviceChannel(deviceID string) (string, error) {
	var d model.Device
	RWLock.RLock()
	ret := DB.
		Where(model.Device{DeviceID: deviceID}).
		Select("channel").
		First(&d)
	RWLock.RUnlock()
	if ret.Error != nil {
		return "", ret.Error
	}
	return d.Channel, nil
}

func (deviceUtil) GetDevice(deviceID string) (model.Device, error) {
	var d model.Device
	RWLock.RLock()
	ret := DB.
		Where(model.Device{DeviceID: deviceID}).
		First(&d)
	RWLock.RUnlock()
	if ret.Error != nil {
		return model.Device{}, ret.Error
	}
	return d, nil
}

func (deviceUtil) CheckDuplicateDeviceToken(deviceID string, token string) (bool, error) {
	var cnt int64
	RWLock.RLock()
	ret := DB.
		Where(model.Device{DeviceID: deviceID, Token: token}).
		Count(&cnt)
	RWLock.RUnlock()
	if ret.Error != nil {
		return false, ret.Error
	}
	return cnt > 0, nil
}

func (deviceUtil) DeleteDevice(deviceID string) error {
	RWLock.Lock()
	ret := DB.
		Where(model.Device{DeviceID: deviceID}).
		Delete(&model.Device{})
	RWLock.Unlock()
	return ret.Error
}

func (deviceUtil) SafeDeleteDevice(userID string, deviceID string) error {
	RWLock.Lock()
	ret := DB.
		Where(model.Device{UserID: userID, DeviceID: deviceID}).
		Delete(&model.Device{})
	RWLock.Unlock()
	return ret.Error
}
