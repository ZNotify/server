package misc

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/ZNotify/server/app/api/common"
	"github.com/ZNotify/server/app/api/middleware"
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
