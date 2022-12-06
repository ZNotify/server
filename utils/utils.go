package utils

import (
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
)

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

func RequireFile(path string) {
	// check if path contains folder
	parts := strings.Split(path, "/")
	if len(parts) > 1 {
		// create folder
		err := os.MkdirAll(strings.Join(parts[:len(parts)-1], "/"), 0755)
		if err != nil {
			zap.S().Fatalf("Failed to create folder: %+v", err)
		}
	}

	va, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			zap.S().Debug("File not exist: ", path)
			_, err := os.Create(path)
			if err != nil {
				zap.S().Fatalf("Failed to create file: %+v", err)
			}
		}
	} else {
		if va.IsDir() {
			zap.S().Fatalf("Path is a directory: %s", path)
		}
		_, err := os.OpenFile(path, os.O_RDWR, 0666)
		if err != nil {
			zap.S().Fatalf("Failed to open file: %+v", err)
		}
	}
}

func TokenClean(token string) string {
	token = strings.Trim(token, " ")
	token = strings.Trim(token, "\t")
	token = strings.Trim(token, "\r")
	token = strings.Trim(token, "\n")
	return token
}
