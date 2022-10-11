package model

import (
	"gorm.io/gorm/clause"
	. "notify-api/db"
	"notify-api/db/entity"
)

type tokenModel struct{}

var TokenUtils = tokenModel{}

// CreateOrUpdate always use the new token
func (_ tokenModel) CreateOrUpdate(userID string, deviceID string, channel string, token string) (entity.PushToken, error) {
	pt := entity.PushToken{
		UserID:   userID,
		DeviceID: deviceID,
		Channel:  channel,
		Token:    token,
	}
	ret := DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "device_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"token", "channel", "updated_at"}),
	}).Create(&pt)
	if ret.Error != nil {
		return entity.PushToken{}, ret.Error
	}
	return pt, nil
}

func (_ tokenModel) GetChannelTokens(userID string, channel string) ([]string, error) {
	var pts []entity.PushToken
	ret := DB.Where("user_id = ? AND channel = ?", userID, channel).Find(&pts)
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
	ret := DB.Where("user_id = ? AND device_id = ?", userID, deviceID).Delete(entity.PushToken{})
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}
