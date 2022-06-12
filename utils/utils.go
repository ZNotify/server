package utils

import (
	"fmt"
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

func IsTestInstance() bool {
	_, ok := os.LookupEnv("TEST")
	return ok
}
