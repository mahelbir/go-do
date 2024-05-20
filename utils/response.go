package utils

import (
	"github.com/gin-gonic/gin"
)

func NewResponse(status bool, message string, data interface{}) gin.H {
	return gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	}
}

func UnauthorizedResponse() gin.H {
	return NewResponse(false, "Unauthorized", nil)
}

func ForbiddenResponse() gin.H {
	return NewResponse(false, "Resource forbidden", nil)
}

func NotFoundResponse() gin.H {
	return NewResponse(false, "Resource not found", nil)
}

func ErrorResponse() gin.H {
	return NewResponse(false, "Something went wrong, please try again later", nil)
}
