package user

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/ZNotify/server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var users []string

// Init read file users.txt to get user list
func Init() {
	if utils.IsTestInstance() {
		users = append(users, "test")
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

// IsUser judge user if in the user list
func IsUser(user string) bool {
	for _, u := range users {
		if u == user {
			return true
		}
	}
	return false
}

func RequireAuth(c *gin.Context) (string, error) {
	userID := c.Param("user_id")
	result := IsUser(userID)
	if !result {
		c.String(http.StatusForbidden, "Unauthorized")
		return "", errors.New("Unauthorized")
	}
	return userID, nil
}
