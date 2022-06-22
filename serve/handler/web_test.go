package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"notify-api/user"
	"testing"
)

func TestWeb(t *testing.T) {
	gin.SetMode(gin.TestMode)
	user.Controller.Init()
	t.Run("check html", func(t *testing.T) {
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
}
