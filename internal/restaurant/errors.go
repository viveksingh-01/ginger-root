package restaurant

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viveksingh-01/ginger-root/internal/httpx"
)

func (h *Handler) badRequest(c *gin.Context, code, message string) {
	h.respondError(c, http.StatusBadRequest, code, message)
}

func (h *Handler) internalError(c *gin.Context, err error) {
	requestID, _ := c.Get("request_id")

	log.Printf(
		"level=error msg=\"internal server error\" error=%v request_id=%v",
		err,
		requestID,
	)

	h.respondError(
		c,
		http.StatusInternalServerError,
		"INTERNAL_ERROR",
		"something went wrong",
	)
}

func (h *Handler) respondError(c *gin.Context, status int, code, msg string) {
	c.JSON(status, httpx.ErrorResponse{Error: code, Message: msg})
}
