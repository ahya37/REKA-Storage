package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	r *gin.RouterGroup,
	handler *Handler,
) {
	r.POST("/login", handler.Login)
}
