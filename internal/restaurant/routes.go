package restaurant

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, handler *Handler) {
	rg.GET("/restaurants", handler.List)
	rg.GET("/restaurants/:restaurantId", handler.GetRestaurant)
	rg.GET("restaurants/search", handler.Search)
}
