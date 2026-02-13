package restaurantmenu

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, handler *Handler) {
	rg.GET("/restaurants/:restaurantId/menu", handler.GetRestaurantMenu)
}
