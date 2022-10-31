package user

import (
	"notify-api/utils/config"
)

var users []string

// Init read file users.txt to get user list
func Init() {
	if config.IsTest() {
		users = append(users, "test")
	} else {
		users = config.Config.Users
	}
}

// Is judge user if in the user list
func Is(user string) bool {
	for _, u := range users {
		if u == user {
			return true
		}
	}
	return false
}

func Users() []string {
	return users
}
