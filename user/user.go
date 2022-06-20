package user

import (
	"bufio"
	"fmt"
	"notify-api/utils"
	"os"
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

	c.users = make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c.users = append(c.users, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
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
