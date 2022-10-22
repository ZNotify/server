package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Duration(c *gin.Context) {
	currentTime := time.Now()
	c.Next()
	duration := time.Since(currentTime)
	c.Writer.Header().Set("X-Duration", duration.String())
}
