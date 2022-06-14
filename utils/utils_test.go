package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBreakOnError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("test BreakOnError with no error", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		BreakOnError(c, nil)
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}
	})

	t.Run("test BreakOnError with error", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("Expected panic, got nil")
			}
			if w.Code != http.StatusInternalServerError {
				t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
			}
		}()
		BreakOnError(c, fmt.Errorf("error"))

	})
}

func TestCheckInternetConnection(t *testing.T) {
	t.Run("test check internet connection", func(t *testing.T) {
		defer func() {
			err := recover()
			if err != nil {
				t.Errorf("Expected no panic, got %v", err)
				t.Errorf("You should keep global network access.")
			}
		}()
		CheckInternetConnection()
	})
}

func TestIsTestInstance(t *testing.T) {
	t.Run("test is test instance", func(t *testing.T) {
		if got := IsTestInstance(); got != true {
			t.Errorf("IsTestInstance() not working.")
		}
	})
}
