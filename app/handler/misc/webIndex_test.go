package misc

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"notify-api/app/common"

	"github.com/gin-gonic/gin"
)

func TestIndex(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("check html", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/", nil)
		common.WrapHandler(WebIndex)(c)
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}
	})
}
