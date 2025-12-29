package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	r *gin.RouterGroup,
	handler *Handler,
) {
	r.GET("/profile", handler.Profile)
	r.POST("/register", handler.Register)
}
