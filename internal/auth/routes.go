package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, handler *Handler) {
	rg.POST("/auth/signup", handler.Signup)
	rg.POST("/auth/login", handler.Login)
}

func RegisterMeHandler(rg *gin.RouterGroup, handler *Handler) {
	rg.GET("/me", handler.GetUser)
}
