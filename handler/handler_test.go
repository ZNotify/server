package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	log.Println("Set Gin to Test Mode")
	ret := m.Run()
	os.Exit(ret)
}
