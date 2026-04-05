package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, handler *Handler) {
	rg.GET("/auth/signup", handler.Signup)
	rg.GET("/auth/login", handler.Login)
}
