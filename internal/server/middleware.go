package server

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/viveksingh-01/ginger-root/internal/auth"
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

const ContextUserID = "userId"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"statusCode":    1,
				"statusMessage": "Authorization header missing",
			})
			c.Abort()
			return
		}

		// Validate format: Bearer <token>
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"statusCode":    1,
				"statusMessage": "Authorization header missing",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Validate token
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"statusCode":    1,
				"statusMessage": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		// Inject userId into context
		c.Set(ContextUserID, claims.UserID)

		// Continue request
		c.Next()
	}
}

// OptionalAuth tries to authenticate a request if an Authorization header is present.
// If token validation succeeds, it injects `userId` into the request context.
func OptionalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.Next()
			return
		}

		// Validate format: Bearer <token>
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.Abort()
			return
		}

		token := parts[1]
		claims, err := auth.ValidateToken(token)

		if err == nil {
			c.Set("userId", claims.UserID)
		}

		c.Next()
	}
}
