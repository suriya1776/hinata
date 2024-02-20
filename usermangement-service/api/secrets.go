// secrets.go

package api

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	once         sync.Once
	randomSecret string
	mu           sync.Mutex
)

// initSecret initializes the random secret key once.
func initSecret() {
	randomSecret = generateRandomSecret()
}

// GetRandomSecret returns the random secret key for JWT signing.
func GetRandomSecret() string {
	once.Do(initSecret)
	return randomSecret
}

// UpdateRandomSecret triggers an update of the random secret key.
func UpdateRandomSecret() {
	mu.Lock()
	defer mu.Unlock()
	randomSecret = generateRandomSecret()
}

// generateRandomSecret generates a random secret key.
func generateRandomSecret() string {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic("Failed to generate random secret key")
	}

	return base64.URLEncoding.EncodeToString(key)
}

// UpdateSecretHandler handles the update of the secret key.
func UpdateSecretHandler(c *gin.Context) {
	// Logic to update the secret key
	UpdateRandomSecret()
	c.JSON(http.StatusOK, gin.H{"message": "Secret key updated successfully"})
}
