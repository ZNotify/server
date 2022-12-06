package user

var users []string

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
