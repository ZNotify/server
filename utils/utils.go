package utils

import (
	"flag"
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
		panic(err)
	}
}

var isTest = -1

func IsTestInstance() bool {
	if isTest != -1 {
		return isTest == 1
	}
	_, ok := os.LookupEnv("TEST")
	if ok {
		isTest = 1
		return true
	}
	f := flag.Lookup("test.v")
	if f != nil {
		isTest = 1
		return true
	}
	isTest = 0
	return false
}
