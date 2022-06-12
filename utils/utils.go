package utils

import (
	"errors"
	"fmt"
	"github.com/ZNotify/server/db"
	"github.com/ZNotify/server/push"
	"github.com/ZNotify/server/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func BreakOnError(c *gin.Context, err error) {
	if err != nil {
		e := c.AbortWithError(500, err)
		if e != nil {
			panic(e)
		}
	}
}

// CheckInternetConnection Check internet connection to google
func CheckInternetConnection() {
	_, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		fmt.Println("No global internet connection")
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}

func RequireAuth(c *gin.Context) (string, error) {
	userID := c.Param("user_id")
	result := user.IsUser(userID)
	if !result {
		c.String(http.StatusForbidden, "Unauthorized")
		return "", errors.New("Unauthorized")
	}
	return userID, nil
}

func IsTestInstance() bool {
	_, ok := os.LookupEnv("TEST")
	return ok
}

func GlobalInit() {
	go CheckInternetConnection()
	db.Init()
	push.Init()
	user.Init()
}
