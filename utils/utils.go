package utils

import (
	"flag"
	"github.com/gin-gonic/gin"
	"os"
	"regexp"
	"sync"
	"time"
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

var uuidRe = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

func IsUUID(uuid string) bool {
	return uuidRe.MatchString(uuid)
}

func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false
	case <-time.After(timeout):
		return true
	}
}
