package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-do/utils"
	"net/http"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, utils.NewResponse(false, "Authorization header is required", nil))
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.UnauthorizedResponse())
			c.Abort()
			return
		}

		c.Set("userID", claims.ID)
		c.Set("isAdmin", claims.IsAdmin)
		c.Set("fullName", claims.FullName)
		c.Set("email", claims.Email)
		c.Next()
	}
}
