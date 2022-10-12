package model

import (
	"time"

	. "notify-api/db"
	"notify-api/db/entity"
)

type messageUtils struct{}

var MessageUtils = messageUtils{}

func (_ messageUtils) Add(id string, userID string, title string, content string, long string, createdAt time.Time) (entity.Message, error) {
	msg := entity.Message{
		ID:        id,
		UserID:    userID,
		Title:     title,
		Content:   content,
		Long:      long,
		CreatedAt: createdAt,
	}
	RWLock.Lock()
	ret := DB.Create(&msg)
	RWLock.Unlock()
	if ret.Error != nil {
		return entity.Message{}, ret.Error
	}
	return msg, nil
}

func (_ messageUtils) Get(id string) (entity.Message, error) {
	var msg entity.Message

	ret := DB.Model(entity.Message{}).Where("id = ?", id).First(&msg)
	if ret.Error != nil {
		return entity.Message{}, ret.Error
	}

	return msg, nil
}

func (_ messageUtils) GetUserMessageAfter(userID string, after time.Time) ([]entity.Message, error) {
	var messages []entity.Message
	ret := DB.Where("user_id = ?", userID).
		Where("created_at > ?", after).
		Order("created_at asc").
		Find(&messages)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return messages, nil
}

func (_ messageUtils) GetMessageInMonth(userID string) ([]entity.Message, error) {
	var messages []entity.Message
	ret := DB.Where("user_id = ?", userID).
		Where("created_at > ?", time.Now().AddDate(0, 0, -30)).
		Order("created_at desc").
		Find(&messages)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return messages, nil
}

func (_ messageUtils) Delete(userID string, msgID string) error {
	ret := DB.Where("user_id = ?", userID).
		Where("id = ?", msgID).
		Delete(&entity.Message{})
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}
