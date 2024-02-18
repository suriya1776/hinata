package api

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/suriya1776/hinata/crm-service/database"
	"github.com/suriya1776/hinata/crm-service/models"
	"golang.org/x/crypto/bcrypt"
)

const tokenExpiration = 1 * time.Minute // Token expiration time
func LoginHandler(c *gin.Context) {
	var loginRequest models.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := database.GetUserByUsername(loginRequest.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := generateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

func generateToken(user *models.BankUser) (string, error) {
	// Create a new JWT token with user information and expiration time
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(tokenExpiration).Unix(),
	})

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte(GetRandomSecret()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
