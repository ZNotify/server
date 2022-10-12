package model

import (
	. "notify-api/db"
	"notify-api/db/entity"
)

type tokenModel struct{}

var TokenUtils = tokenModel{}

// CreateOrUpdate always use the new token
func (_ tokenModel) CreateOrUpdate(userID string, deviceID string, channel string, token string) error {
	ptn := map[string]interface{}{
		"channel": channel,
		"token":   token,
	}
	var pt entity.PushToken
	RWLock.Lock()
	ret := DB.
		Where(entity.PushToken{UserID: userID, DeviceID: deviceID}).
		Assign(ptn).
		FirstOrCreate(&pt)
	RWLock.Unlock()

	return ret.Error
}

func (_ tokenModel) GetChannelTokens(userID string, channel string) ([]string, error) {
	var pts []entity.PushToken
	ret := DB.Where(&entity.PushToken{
		UserID:  userID,
		Channel: channel,
	}).Find(&pts)
	if ret.Error != nil {
		return nil, ret.Error
	}
	var tokens []string
	for _, pt := range pts {
		tokens = append(tokens, pt.Token)
	}
	return tokens, nil
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
