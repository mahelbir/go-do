package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Auth middleware
		fmt.Println("Auth middleware")
		c.Set("testKey", "auth middleware")
		c.Next()
	}
}
