package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Signup(c *gin.Context) {
	var req *SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			StatusCode:    1,
			StatusMessage: err.Error(),
		})
		return
	}

	user, err := h.service.Signup(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			StatusCode:    1,
			StatusMessage: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Response{
		StatusCode:    0,
		StatusMessage: "done successfully",
		Data: &AuthResponse{
			User:  ToUserResponse(user),
			Token: "",
		},
	})
}
