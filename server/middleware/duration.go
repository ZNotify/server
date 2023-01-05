package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type DurationWriter struct {
	gin.ResponseWriter
	startTime time.Time
}

func (w DurationWriter) WriteHeader(code int) {
	w.Header().Set("X-Handle-Time", strconv.FormatInt(time.Since(w.startTime).Microseconds(), 10))
	w.ResponseWriter.WriteHeader(code)
}

func Duration(c *gin.Context) {
	currentTime := time.Now()
	c.Writer = &DurationWriter{c.Writer, currentTime}
	c.Next()
}
