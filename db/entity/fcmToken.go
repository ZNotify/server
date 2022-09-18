package entity

import "time"

type FCMTokens struct {
	ID             string
	UserID         string
	CreatedAt      time.Time
	RegistrationID string
}
