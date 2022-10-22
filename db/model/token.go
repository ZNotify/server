package model

import (
	"time"

	. "notify-api/db"
	"notify-api/db/entity"
)

type tokenModel struct{}

var TokenUtils = tokenModel{}

// CreateOrUpdate always use the new token
func (_ tokenModel) CreateOrUpdate(userID string, deviceID string, channel string, token string) error {
	var pt entity.PushToken
	RWLock.RLock()
	ret := DB.
		Where(entity.PushToken{UserID: userID, DeviceID: deviceID}).
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

	RWLock.Lock()
	ret = DB.Save(&pt)
	RWLock.Unlock()

	return ret.Error
}

func (_ tokenModel) GetChannelTokens(userID string, channel string) ([]string, error) {
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

func (_ tokenModel) GetDeviceToken(userID string, deviceID string) (entity.PushToken, error) {
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

func (_ tokenModel) Delete(userID string, deviceID string) error {
	RWLock.Lock()
	ret := DB.Delete(entity.PushToken{
		UserID:   userID,
		DeviceID: deviceID,
	})
	RWLock.Unlock()
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}
