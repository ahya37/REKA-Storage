package response

import "github.com/gin-gonic/gin"

type Meta struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type APIResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, APIResponse{
		Meta: Meta{
			Code:    code,
			Status:  "success",
			Message: message,
		},
		Data: data,
	})
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(code, APIResponse{
		Meta: Meta{
			Code:    code,
			Status:  "error",
			Message: message,
		},
	})
}
