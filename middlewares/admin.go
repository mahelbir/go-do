package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-do/utils"
	"net/http"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		isAdmin := c.GetBool("isAdmin")
		if !isAdmin {
			c.JSON(http.StatusForbidden, utils.ForbiddenResponse())
			c.Abort()
			return
		}

		c.Next()

	}
}
