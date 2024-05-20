package controllers

import (
	"github.com/gin-gonic/gin"
	"go-do/models"
	"go-do/services"
	"go-do/utils"
	"net/http"
)

func GetUser(c *gin.Context) {
	user, err := services.UserSvc.GetByID(c.GetInt("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, utils.NotFoundResponse())
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(true, "User found", models.ToPublicUser(user)))
}
