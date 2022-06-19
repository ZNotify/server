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

func (t fcmTokensUtils) Add(userID string, regID string) (FCMTokens, error) {
	token := FCMTokens{
		ID:             uuid.New().String(),
		UserID:         userID,
		RegistrationID: regID,
	}
	ret := db.DB.Create(&token)
	if ret.Error != nil {
		return FCMTokens{}, ret.Error
	}
	return token, nil
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
