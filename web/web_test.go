package web

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestWebHandler(t *testing.T) {
	t.Run("test web handler", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/", nil)
		r := gin.New()
		r.StaticFS("/fs", StaticHttpFS)

		req := httptest.NewRequest("GET", "/fs/robots.txt", nil)
		r.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}

		if !strings.Contains(w.Header().Get("Content-Type"), "text/plain") {
			t.Errorf("Expected content type %s, got %s", "text/plain", w.Header().Get("Content-Type"))
		}
	})
}
