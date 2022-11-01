package model

import (
	"time"

	. "notify-api/db"
	"notify-api/db/entity"
)

type tokenModel struct{}

var TokenUtils = tokenModel{}

// CreateOrUpdate always use the new token
func (tokenModel) CreateOrUpdate(userID string, deviceID string, channel string, token string) error {
	var pt entity.PushToken
	RWLock.RLock()
	ret := DB.
		Where(entity.PushToken{DeviceID: deviceID}).
		FirstOrInit(&pt)
	RWLock.RUnlock()
	if ret.Error != nil {
		return ret.Error
	}

	// FIXME: use register mechanism to process token
	if channel == "WebSocketHost" && token == "" {
		if pt.Channel != "WebSocketHost" {
			pt.Token = time.Now().Format(time.RFC3339Nano)
		}
	} else {
		pt.Token = token
	}
	pt.Channel = channel
	pt.UserID = userID

	RWLock.Lock()
	ret = DB.Save(&pt)
	RWLock.Unlock()

	return ret.Error
}

func (tokenModel) GetUserChannelTokens(userID string, channel string) ([]string, error) {
	var pts []entity.PushToken
	RWLock.RLock()
	ret := DB.Where(&entity.PushToken{
		UserID:  userID,
		Channel: channel,
	}).Find(&pts)
	RWLock.RUnlock()
	if ret.Error != nil {
		return nil, ret.Error
	}
	var tokens []string
	for _, pt := range pts {
		tokens = append(tokens, pt.Token)
	}
	return tokens, nil
}

func (tokenModel) GetUserDeviceToken(userID string, deviceID string) (entity.PushToken, error) {
	var pt entity.PushToken
	RWLock.RLock()
	ret := DB.Where(&entity.PushToken{
		UserID:   userID,
		DeviceID: deviceID,
	}).First(&pt)
	RWLock.RUnlock()
	if ret.Error != nil {
		return entity.PushToken{}, ret.Error
	}
	return pt, nil
}

func (tokenModel) GetDeviceToken(deviceID string) (entity.PushToken, error) {
	var pt entity.PushToken
	RWLock.RLock()
	ret := DB.Where(&entity.PushToken{
		DeviceID: deviceID,
	}).First(&pt)
	RWLock.RUnlock()
	if ret.Error != nil {
		return entity.PushToken{}, ret.Error
	}
	return pt, nil
}

func (tokenModel) Delete(deviceID string) error {
	RWLock.Lock()
	ret := DB.
		Where(&entity.PushToken{DeviceID: deviceID}).
		Delete(entity.PushToken{})
	RWLock.Unlock()
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}
