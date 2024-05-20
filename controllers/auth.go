package controllers

import (
	"github.com/gin-gonic/gin"
	"go-do/models"
	"go-do/services"
	"go-do/utils"
	"net/http"
)

func Login(c *gin.Context) {
	var details models.User
	if err := c.ShouldBindJSON(&details); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(false, "Invalid request", nil))
		return
	}

	if details.Email == "" || details.Password == "" {
		c.JSON(http.StatusBadRequest, utils.NewResponse(false, "Email and password are required", nil))
		return
	}

	user, err := services.UserSvc.GetByEmail(details.Email)
	if err != nil || !utils.VerifyPassword(user.Password, details.Password) {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(false, "Invalid email or password", nil))
		return
	}

	tokenString, err := utils.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse())
		return
	}

	type LoginResponse struct {
		Email    string `json:"email"`
		FullName string `json:"full_name"`
		Token    string `json:"token"`
	}

	c.JSON(http.StatusOK, utils.NewResponse(true, "Login successful", &LoginResponse{
		Email:    user.Email,
		FullName: user.FullName,
		Token:    tokenString,
	}))
}
