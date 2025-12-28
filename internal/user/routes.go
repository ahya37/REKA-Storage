package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	r *gin.RouterGroup,
	handler *Handler,
) {
	r.POST("/profile", handler.Profile)
}
