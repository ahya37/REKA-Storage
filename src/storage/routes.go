package storage

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	r *gin.RouterGroup,
	handler *Handler,
) {
	r.POST("/upload", handler.Upload)
	r.GET("/list", handler.ListByUser)
}
