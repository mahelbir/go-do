package controllers

import (
	"github.com/gin-gonic/gin"
	"go-do/services"
	"go-do/utils"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(false, "Invalid user ID", nil))
		return
	}

	user, err := services.UserSvc.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.NotFoundResponse())
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(true, "User found", user))
}
