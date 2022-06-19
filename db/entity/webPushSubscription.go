package entity

import (
	"github.com/google/uuid"
	"notify-api/db"
	"time"
)

type WebSubscription struct {
	ID           string
	UserID       string
	CreatedAt    time.Time
	Subscription string
}

func (s WebSubscription) Add() {
	db.DB.Create(s)
}

type webSubModel struct{}

var WebSubUtils = webSubModel{}

func (_ webSubModel) Add(userID string, sub string) error {
	ret := db.DB.Create(&WebSubscription{
		ID:           uuid.New().String(),
		UserID:       userID,
		Subscription: sub,
	})
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}

func (_ webSubModel) Count(userID string, sub string) int64 {
	var cnt int64
	db.DB.Model(&WebSubscription{}).
		Where("user_id = ?", userID).
		Where("subscription = ?", sub).
		Count(&cnt)
	return cnt
}

func (_ webSubModel) Get(userID string) []WebSubscription {
	var subs []WebSubscription
	db.DB.Where("user_id = ?", userID).Find(&subs)
	return subs
}
