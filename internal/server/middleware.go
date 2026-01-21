package server

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		requestID := uuid.New().String()
		c.Set("request_id", requestID)
		c.Writer.Header().Set("X-Request-ID", requestID)

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		log.Printf(
			"level=info method=%s path=%s status=%d latency=%s request_id=%s",
			c.Request.Method,
			c.Request.URL.Path,
			status,
			latency,
			requestID,
		)
	}
}
