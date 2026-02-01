package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create router
	r := gin.New()

	// Register handler
	healthHandler := NewHandler()
	r.GET("/health", healthHandler.Check)

	// Create fake HTTP request
	req, err := http.NewRequest(http.MethodGet, "/health", nil)
	assert.NoError(t, err)

	// Record the response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert response body
	expectedBody := `{"status": "OK"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}
