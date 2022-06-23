package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWebHandler(t *testing.T) {
	Init()
	t.Run("test web handler", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/", nil)
		r := gin.New()
		r.StaticFS("/fs", StaticHttpFS)

		req := httptest.NewRequest("GET", "/fs/favicon.ico", nil)
		r.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}
		if w.Header().Get("Content-Type") != "image/x-icon" {
			t.Errorf("Expected content type %s, got %s", "image/x-icon", w.Header().Get("Content-Type"))
		}
	})
}
