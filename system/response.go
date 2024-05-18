package system

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  bool
	Message string
	Data    any
}

func NewResponse(response Response) gin.H {
	return gin.H{
		"Status":  response.Status,
		"Message": response.Message,
		"Data":    response.Data,
	}
}

func UnauthorizedResponse() gin.H {
	return NewResponse(Response{Message: "Unauthorized"})
}

func ForbiddenResponse() gin.H {
	return NewResponse(Response{Message: "Resource forbidden"})
}

func NotFoundResponse() gin.H {
	return NewResponse(Response{Message: "Resource not found"})
}

func ErrorResponse() gin.H {
	return NewResponse(Response{Message: "Something went wrong, please try again later"})
}
