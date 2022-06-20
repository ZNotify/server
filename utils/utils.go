package utils

import (
	"flag"
	"github.com/gin-gonic/gin"
	"os"
)

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
	if gin.Mode() == gin.TestMode {
		isTest = 1
		return true
	}
	isTest = 0
	return false
}
