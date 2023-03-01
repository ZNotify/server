package utils

import (
	"encoding/json"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

func IsUUID(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
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

func YamlStringClean(token string) string {
	token = strings.Trim(token, " ")
	token = strings.Trim(token, "\t")
	token = strings.Trim(token, "\r")
	token = strings.Trim(token, "\n")
	return token
}

func OAuthTokenSerialize(token *oauth2.Token) string {
	if d, err := json.Marshal(token); err == nil {
		return string(d)
	}
	return ""
}

func OAuthTokenDeserialize(token string) *oauth2.Token {
	var t oauth2.Token
	if err := json.Unmarshal([]byte(token), &t); err == nil {
		return &t
	}
	return nil
}
