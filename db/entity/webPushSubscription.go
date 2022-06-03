package entity

import "time"

type WebSubscription struct {
	ID           string
	UserID       string
	CreatedAt    time.Time
	Subscription string
}
