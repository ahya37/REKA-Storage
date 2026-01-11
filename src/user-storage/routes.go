package userstorage

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	r *gin.RouterGroup,
	handler *Handler,
) {
	r.GET("/usage", handler.GetUsage)
}
