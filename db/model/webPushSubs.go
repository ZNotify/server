package model

import (
	"github.com/google/uuid"
	. "notify-api/db"
	"notify-api/db/entity"
)

type webSubUtils struct{}

var WebSubUtils = webSubUtils{}

func (_ webSubUtils) Add(userID string, sub string) (entity.WebSubscription, error) {
	s := entity.WebSubscription{
		ID:           uuid.New().String(),
		UserID:       userID,
		Subscription: sub,
	}
	ret := DB.Create(&s)
	if ret.Error != nil {
		return entity.WebSubscription{}, ret.Error
	}
	return s, nil
}

func (_ webSubUtils) Count(userID string, sub string) int64 {
	var cnt int64
	DB.Model(&entity.WebSubscription{}).
		Where("user_id = ?", userID).
		Where("subscription = ?", sub).
		Count(&cnt)
	return cnt
}

func (_ webSubUtils) Get(userID string) ([]entity.WebSubscription, error) {
	var subs []entity.WebSubscription
	ret := DB.Where("user_id = ?", userID).Find(&subs)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return subs, nil
}
