package entity

import (
	"github.com/google/uuid"
	"notify-api/db"
	"time"
)

type FCMTokens struct {
	ID             string
	UserID         string
	CreatedAt      time.Time
	RegistrationID string
}

var FCMUtils = fcmTokensUtils{}

type fcmTokensUtils struct{}

func (t fcmTokensUtils) Add(userID string, regID string) error {
	ret := db.DB.Create(FCMTokens{
		ID:             uuid.New().String(),
		UserID:         userID,
		RegistrationID: regID,
	})
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}

func (_ fcmTokensUtils) Count(userID string, regID string) (int64, error) {
	var cnt int64
	ret := db.DB.Model(&FCMTokens{}).
		Where("user_id = ?", userID).
		Where("registration_id = ?", regID).
		Count(&cnt)
	if ret.Error != nil {
		return 0, ret.Error
	}
	return cnt, nil
}

func (_ fcmTokensUtils) Get(userID string) []FCMTokens {
	var tokens []FCMTokens
	db.DB.Where("user_id = ?", userID).Find(&tokens)
	return tokens
}
