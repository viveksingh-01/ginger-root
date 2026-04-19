package cart

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, handler *Handler) {
	rg.POST("/cart", handler.AddToCart)
	rg.GET("/cart", handler.GetCart)
}
