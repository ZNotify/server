package util

import (
	"time"

	. "notify-api/db"
	"notify-api/db/model"
	entity2 "notify-api/push/entity"
)

type messageUtil struct{}

var MessageUtil = messageUtil{}

func (messageUtil) Add(msgID string, userID string, title string, content string, long string, priority entity2.Priority) (model.Message, error) {
	msg := model.Message{
		MessageID: msgID,
		UserID:    userID,
		Title:     title,
		Content:   content,
		Long:      long,
		Priority:  priority,
		CreatedAt: time.Now(),
	}
	RWLock.Lock()
	ret := DB.Create(&msg)
	RWLock.Unlock()
	if ret.Error != nil {
		return model.Message{}, ret.Error
	}
	return msg, nil
}

func (messageUtil) Get(id string) (model.Message, error) {
	var msg model.Message

	RWLock.RLock()
	ret := DB.Model(model.Message{}).Where("id = ?", id).First(&msg)
	RWLock.RUnlock()
	if ret.Error != nil {
		return model.Message{}, ret.Error
	}

	return msg, nil
}

func (messageUtil) GetUserMessageAfter(userID string, after time.Time) ([]model.Message, error) {
	var messages []model.Message
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

func (messageUtil) GetMessageInMonth(userID string) ([]model.Message, error) {
	var messages []model.Message
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

func (messageUtil) Delete(userID string, msgID string) error {
	RWLock.Lock()
	ret := DB.Where("user_id = ?", userID).
		Where("id = ?", msgID).
		Delete(&model.Message{})
	RWLock.Unlock()
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}
