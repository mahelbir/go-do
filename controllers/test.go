package controllers

import (
	"github.com/gin-gonic/gin"
	"go-do/utils"
	"net/http"
)

func TestController(c *gin.Context) {
	c.JSON(http.StatusOK, utils.NewResponse(true, "Test controller", nil))
}
