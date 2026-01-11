package userstorage

import (
	"net/http"
	"reka-storage/src/shared/response"
	"reka-storage/src/user-storage/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *services.UserStorageService
}

func NewUserStorageHandler(service *services.UserStorageService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetUsage(c *gin.Context) {
	userID := c.MustGet("userID").(string)

	usage, err := h.service.GetUsage(
		c.Request.Context(),
		userID,
	)

	if err != nil {
		response.Error(
			c,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	response.Success(
		c,
		http.StatusCreated,
		"user storage usage retrieved successfully",
		usage,
	)
}
