package user

import (
	"bufio"
	"os"

	"go.uber.org/zap"

	"notify-api/utils"
)

type userController struct {
	users []string
}

var Controller = userController{
	users: make([]string, 0),
}

// Init read file users.txt to get user list
func (c *userController) Init() {
	if utils.IsTestInstance() {
		c.users = append(c.users, "test")
		return
	}

	file, err := os.Open("data/users.txt")
	if err != nil {
		zap.S().Fatalf("Failed to open users.txt: %+v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			zap.S().Fatalf("Failed to close users.txt: %+v", err)
		}
	}(file)

	c.users = make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c.users = append(c.users, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		zap.S().Fatalf("Failed to read users.txt: %+v", err)
	}
}

// Is judge user if in the user list
func (c *userController) Is(user string) bool {
	for _, u := range c.users {
		if u == user {
			return true
		}
	}
	return false
}

func (c *userController) Users() []string {
	return c.users
}
