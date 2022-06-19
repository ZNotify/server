package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"notify-api/user"
	"testing"
)

func TestCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)
	user.Init()
	t.Run("check success", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/check?user_id=test", nil)
		Check(c)
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}
		if w.Body.String() != "true" {
			t.Errorf("Expected body %s, got %s", "true", w.Body.String())
		}
	})

	t.Run("check fail", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/check?user_id=error", nil)
		Check(c)
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}
		if w.Body.String() != "false" {
			t.Errorf("Expected body %s, got %s", "false", w.Body.String())
		}
	})
}
