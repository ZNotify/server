package user

import (
	"bufio"
	"fmt"
	"os"
)

// Init read file users.txt to get user list
func Init() {
	file, err := os.Open("users.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}(file)

	users = make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		users = append(users, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var users []string

// IsUser judge user if in the user list
func IsUser(user string) bool {
	for _, u := range users {
		if u == user {
			return true
		}
	}
	return false
}
