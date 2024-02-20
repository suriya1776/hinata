package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/suriya1776/hinata/usermangement-service/database"
	"github.com/suriya1776/hinata/usermangement-service/models"
	"golang.org/x/crypto/bcrypt"
)

// Token will expire in 5 minutes
// const tokenExpiration = 5 * time.Minute // Token expiration time

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

	// Check if the password has expired
	if time.Since(user.LastPasswordUpdate).Minutes() > 2 {
		// Password has expired
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password has expired. Please reset your password"})
		return
	}

	// Continue with regular login logic
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
	expirationTime := time.Now().Add(1 * time.Hour) // Set the expiration time (e.g., one hour from now)

	rolesClaim := jwt.MapClaims{
		"username": user.Username,
		"role":     user.Roles[0],         // Assuming you want to include the first role as a separate field,
		"exp":      expirationTime.Unix(), // Set the expiration time in Unix timestamp format
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, rolesClaim)

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte(GetRandomSecret()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// user login
func UserProfileHandler(c *gin.Context) {
	// Extract user information from the token
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing token"})
		return
	}

	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing token"})
		return
	}

	// Customize this part based on your user data retrieval logic
	userProfile := map[string]interface{}{
		"username": username,
		"role":     role,
		// Add other user details as needed
	}

	c.JSON(http.StatusOK, gin.H{"userProfile": userProfile})
}

// Define the password reset handler
// Define the password reset handler
// Define the password reset handler
func ResetPasswordHandler(c *gin.Context) {
	// Retrieve the necessary information from the request
	var resetRequest models.ResetRequest
	if err := c.ShouldBindJSON(&resetRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the provided old password
	user, err := database.GetUserByUsername(resetRequest.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	fmt.Println("Before update:")
	fmt.Println("LastPasswordUpdate:", user.LastPasswordUpdate)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(resetRequest.OldPassword))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid old password"})
		return
	}

	// Check if the new password and confirm password match
	if resetRequest.NewPassword != resetRequest.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "New password and confirm password do not match"})
		return
	}

	// Update the user's password and reset flag
	hashedNewPassword, err := database.HashPassword(resetRequest.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash the new password"})
		return
	}

	// Update user fields
	user.Password = hashedNewPassword

	// Update the last password update time
	user.LastPasswordUpdate = time.Now()

	fmt.Println("After update:")
	fmt.Println("LastPasswordUpdate:", user.LastPasswordUpdate)

	// Update the user in the database
	err = database.UpdateUser(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successful"})
}
