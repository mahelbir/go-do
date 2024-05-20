package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"go-do/models"
	"os"
	"time"
)

type Claims struct {
	ID       int    `json:"id"`
	IsAdmin  bool   `json:"is_admin"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateJWT(user *models.User) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24)
	claims := &Claims{
		ID:       user.ID,
		IsAdmin:  user.IsAdmin,
		FullName: user.FullName,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
