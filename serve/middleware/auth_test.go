package middleware

import (
	"net/http"
	"net/http/httptest"
	"notify-api/utils/user"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type AuthMiddlewareTestSuite struct {
	suite.Suite
	router   *gin.Engine
	recorder *httptest.ResponseRecorder
}

func (suite *AuthMiddlewareTestSuite) SetupTest() {
	suite.router = gin.New()
	suite.recorder = httptest.NewRecorder()
}

func (suite *AuthMiddlewareTestSuite) SetupSuite() {
	user.Controller.Init()
	gin.SetMode(gin.TestMode)
}

func (suite *AuthMiddlewareTestSuite) TestAuthPassed() {
	suite.router.GET("/:user_id/ok", UserAuth, func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})
	req := httptest.NewRequest("GET", "/test/ok", nil)
	suite.router.ServeHTTP(suite.recorder, req)
	suite.Equal(http.StatusOK, suite.recorder.Code)
}

func (suite *AuthMiddlewareTestSuite) TestAuthFailed() {
	suite.router.GET("/:user_id/ok", UserAuth, func(c *gin.Context) {
		suite.FailNow("Not authed, should not be called")
	})
	req := httptest.NewRequest("GET", "/error/ok", nil)
	suite.router.ServeHTTP(suite.recorder, req)
	suite.Equal(http.StatusUnauthorized, suite.recorder.Code)
}

func TestAuth(t *testing.T) {
	suite.Run(t, new(AuthMiddlewareTestSuite))
}
