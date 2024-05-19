package controllers

import (
	"github.com/gin-gonic/gin"
	"go-do/models"
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

	c.JSON(http.StatusOK, utils.NewResponse(true, "User found", models.ToPublicUser(*user)))
}

func Login(c *gin.Context) {
	var loginDetails models.User
	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(false, "Invalid request", nil))
		return
	}

	user, err := services.UserSvc.GetByEmail(loginDetails.Email)
	if err != nil || !utils.VerifyPassword(user.Password, loginDetails.Password) {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(false, "Invalid email or password", nil))
		return
	}

	tokenString, err := utils.GenerateJWT(user.ID, user.IsAdmin, user.FullName, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse())
		return
	}

	type LoginResponse struct {
		Email    string `json:"email"`
		FullName string `json:"full_name"`
		Token    string `json:"token"`
	}

	c.JSON(http.StatusOK, utils.NewResponse(true, "Login successful", LoginResponse{
		Email:    user.Email,
		FullName: user.FullName,
		Token:    tokenString,
	}))
}
