package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// AppConfig holds application-wide configurations
type AppConfig struct {
	ServerPort string
	DatabaseURI string
}

// LoadConfig loads configurations from environment variables or a `.env` file
func LoadConfig() *AppConfig {
	// Load environment variables from a .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	// Retrieve configuration values
	return &AppConfig{
		ServerPort: getEnv("SERVER_PORT", "9090"),
		DatabaseURI: getEnv("DATABASE_URI", "mongodb://localhost:27017"),
	}
}

// getEnv retrieves environment variables or a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
