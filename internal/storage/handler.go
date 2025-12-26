package storage

import (
	"net/http"
	"reka-storage/internal/shared/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	folder := c.PostForm("folder")

	if err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"File is required",
		)
		return
	}

	if folder == "" {
		response.Error(
			c,
			http.StatusBadRequest,
			"Folder is required",
		)
		return
	}

	userID := c.MustGet("userID").(string)

	result, err := h.service.Upload(c.Request.Context(), userID, file, folder)
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
		http.StatusOK,
		"File uploaded successfully",
		result,
	)
}
