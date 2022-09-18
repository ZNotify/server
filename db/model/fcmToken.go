package model

import (
	"github.com/google/uuid"
	"notify-api/db"
	. "notify-api/db/entity"
)

var FCMTokenUtils = fcmTokensUtils{}

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

func (_ fcmTokensUtils) Get(userID string) ([]FCMTokens, error) {
	var tokens []FCMTokens
	ret := db.DB.Where("user_id = ?", userID).Find(&tokens)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return tokens, nil
}
