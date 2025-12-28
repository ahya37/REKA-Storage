package user

import (
	"net/http"
	"reka-storage/src/shared/response"
	"reka-storage/src/user/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Profile(c *gin.Context) {
	// ambil userID dari middleware
	userID := c.MustGet("userID").(string)

	profile, err := h.service.GetProfile(c.Request.Context(), userID)
	if err != nil {
		response.Error(c, http.StatusNotFound, "user not found")
		return
	}

	response.Success(c, http.StatusOK, "profile retrieved successfully", profile)
}
