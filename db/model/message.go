package model

import (
	"time"

	. "notify-api/db"
	"notify-api/db/entity"
)

type messageUtils struct{}

var MessageUtils = messageUtils{}

func (messageUtils) Add(id string, userID string, title string, content string, long string) (entity.Message, error) {
	// a trick to generate different timestamp for different message
	// FIXME: use an increasing counter to generate different id
	msg := entity.Message{
		ID:        id,
		UserID:    userID,
		Title:     title,
		Content:   content,
		Long:      long,
		CreatedAt: time.Now().Add(time.Nanosecond * 10),
	}
	RWLock.Lock()
	ret := DB.Create(&msg)
	RWLock.Unlock()
	if ret.Error != nil {
		return entity.Message{}, ret.Error
	}
	return msg, nil
}

func (messageUtils) Get(id string) (entity.Message, error) {
	var msg entity.Message

	RWLock.RLock()
	ret := DB.Model(entity.Message{}).Where("id = ?", id).First(&msg)
	RWLock.RUnlock()
	if ret.Error != nil {
		return entity.Message{}, ret.Error
	}

	return msg, nil
}

func (messageUtils) GetUserMessageAfter(userID string, after time.Time) ([]entity.Message, error) {
	var messages []entity.Message
	RWLock.RLock()
	ret := DB.Where("user_id = ?", userID).
		Where("created_at > ?", after).
		Order("created_at asc").
		Find(&messages)
	RWLock.RUnlock()
	if ret.Error != nil {
		return nil, ret.Error
	}
	return messages, nil
}

func (messageUtils) GetMessageInMonth(userID string) ([]entity.Message, error) {
	var messages []entity.Message
	RWLock.RLock()
	ret := DB.Where("user_id = ?", userID).
		Where("created_at > ?", time.Now().AddDate(0, 0, -30)).
		Order("created_at desc").
		Find(&messages)
	RWLock.RUnlock()
	if ret.Error != nil {
		return nil, ret.Error
	}
	return messages, nil
}

func (messageUtils) Delete(userID string, msgID string) error {
	RWLock.Lock()
	ret := DB.Where("user_id = ?", userID).
		Where("id = ?", msgID).
		Delete(&entity.Message{})
	RWLock.Unlock()
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}
