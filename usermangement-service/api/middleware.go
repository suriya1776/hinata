package api

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
			c.Abort()
			return
		}

		// Extract the token without the "Bearer " prefix
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// You should have a secure secret key for signing and validating the token
			return []byte(GetRandomSecret()), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired authorization token"})
			c.Abort()
			return
		}

		// Extract user information from the token claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Check if "username" and "role" keys are present and have the correct type
		username, usernameOK := claims["username"].(string)
		role, roleOK := claims["role"].(string)
		if !usernameOK || !roleOK {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing username or role in token"})
			c.Abort()
			return
		}

		// Set user information in the request context
		c.Set("username", username)
		c.Set("role", role)

		// Continue to the next middleware or handler
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			c.Abort()
			return
		}
		c.Next()
	}
}
