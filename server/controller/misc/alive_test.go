package misc

import (
	"net/http"
	"net/http/httptest"
	"notify-api/server/middleware"
	"testing"

	"github.com/gin-gonic/gin"

	"notify-api/server/types"
)

func TestAlive(t *testing.T) {
	t.Run("test alive handler", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		middleware.ServerTiming(c)
		types.WrapHandler(Alive)(c)
		if w.Code != http.StatusNoContent {
			t.Errorf("Expected status code %d, got %d", http.StatusNoContent, w.Code)
		}
	})
}
