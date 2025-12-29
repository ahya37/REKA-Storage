package user

import (
	"net/http"
	"reka-storage/src/shared/response"
	"reka-storage/src/user/dtos"
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

	user, err := h.service.GetProfile(c.Request.Context(), userID)
	if err != nil {
		response.Error(c, http.StatusNotFound, "user not found")
		return
	}

	profile := dtos.FindByIDDto{
		ID:        user.Id.Hex(),
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}

	response.Success(c, http.StatusOK, "profile retrieved successfully", profile)
}

func (h *Handler) Register(c *gin.Context) {
	var req dtos.RegisterRequestDto

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	user, err := h.service.Register(
		c.Request.Context(),
		req.Username,
		req.Email,
		req.Password,
	)
	if err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	response.Success(
		c,
		http.StatusCreated,
		"Register successfully",
		user,
	)
}
