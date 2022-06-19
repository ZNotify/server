package entity

import (
	"notify-api/db"
	"time"
)

type Message struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Long      string    `json:"long"`
	CreatedAt time.Time `json:"created_at"`
}

type messageModel struct{}

var MessageUtils = messageModel{}

func (_ messageModel) Add(id string, userID string, title string, content string, long string, createdAt time.Time) error {
	ret := db.DB.Create(&Message{
		ID:        id,
		UserID:    userID,
		Title:     title,
		Content:   content,
		Long:      long,
		CreatedAt: createdAt,
	})
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}

func (_ messageModel) GetMessageInMonth(userID string) ([]Message, error) {
	var messages []Message
	ret := db.DB.Where("user_id = ?", userID).
		Where("created_at > ?", time.Now().AddDate(0, 0, -30)).
		Order("created_at desc").
		Find(&messages)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return messages, nil
}

func (_ messageModel) Delete(userID string, msgID string) error {
	ret := db.DB.Where("user_id = ?", userID).
		Where("id = ?", msgID).
		Delete(&Message{})
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}
