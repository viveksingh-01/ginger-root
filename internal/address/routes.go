package address

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, handler *Handler) {
	rg.GET("/addresses", handler.GetAddresses)
	rg.POST("/address", handler.SaveAddress)
}
