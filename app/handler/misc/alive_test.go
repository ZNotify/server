package misc

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"notify-api/app/common"
	"notify-api/app/middleware"

	"github.com/gin-gonic/gin"
)

func TestAlive(t *testing.T) {
	t.Run("test alive handler", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		middleware.ServerTiming(c)
		common.WrapHandler(Alive)(c)
		if w.Code != http.StatusNoContent {
			t.Errorf("Expected status code %d, got %d", http.StatusNoContent, w.Code)
		}
	})
}
