package order

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, handler *Handler) {
	rg.POST("/order", handler.PlaceOrder)
	rg.GET("/order/:orderId/status", handler.TrackOrderStatus)
}
