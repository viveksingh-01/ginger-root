package address

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

func (h *Handler) SaveAddress(c *gin.Context) {
	var req CreateAddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			StatusCode:    1,
			StatusMessage: err.Error(),
		})
		return
	}

	userId := c.GetString("userId")
	address, err := h.service.CreateAddress(c, &req, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			StatusCode:    1,
			StatusMessage: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &Response{
		StatusCode:    0,
		StatusMessage: "Address saved successfully",
		Data:          []AddressResponse{*ToAddressResponse(address)},
	})
}

func (h *Handler) GetAddresses(c *gin.Context) {
	userId := c.GetString("userId")
	addresses, err := h.service.GetAddresses(c, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			StatusCode:    1,
			StatusMessage: err.Error(),
		})
		return
	}

	var resp []AddressResponse
	for _, addr := range addresses {
		resp = append(resp, *ToAddressResponse(addr))
	}

	c.JSON(http.StatusOK, &Response{
		StatusCode:    0,
		StatusMessage: "Addresses fetched successfully",
		Data:          resp,
	})
}

func (h *Handler) DeleteAddress(c *gin.Context) {
	addressId := c.Param("addressId")
	err := h.service.DeleteAddress(c, addressId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			StatusCode:    1,
			StatusMessage: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		StatusCode:    0,
		StatusMessage: "Address deleted successfully",
	})
}
