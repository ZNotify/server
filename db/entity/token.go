package entity

import (
	"gorm.io/gorm/clause"
	. "notify-api/db"
	"time"
)

type PushToken struct {
	ID        string
	UserID    string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeviceID  string
	Channel   string
	Token     string
}

type PushTokenModel struct{}

var PushTokenUtils = PushTokenModel{}

// CreateOrUpdate always use the new token
func (_ PushTokenModel) CreateOrUpdate(userID string, deviceID string, channel string, token string) (PushToken, error) {
	pt := PushToken{
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
		return PushToken{}, ret.Error
	}
	return pt, nil
}
