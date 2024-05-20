package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-do/services"
	"go-do/utils"
	"net/http"
)

func TodoListAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		todoList, err := services.TodoListSvc.GetByID(c.GetInt("id"))

		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse())
			c.Abort()
			return
		}

		if todoList.UserID != c.GetInt("userID") && !c.GetBool("isAdmin") {
			c.JSON(http.StatusForbidden, utils.ForbiddenResponse())
			c.Abort()
			return
		}

		if todoList.ID <= 0 {
			c.JSON(http.StatusNotFound, utils.NotFoundResponse())
			c.Abort()
			return
		}

		c.Set("todoList", todoList)
		c.Next()
	}
}
