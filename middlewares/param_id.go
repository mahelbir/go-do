package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-do/utils"
	"net/http"
	"strconv"
)

func ParamID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(false, "Invalid ID", nil))
			c.Abort()
			return
		}

		c.Set("id", id)
		c.Next()
	}
}
