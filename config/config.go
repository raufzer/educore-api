package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort  string
	DatabaseURI string
}

func LoadConfig() *AppConfig {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Warning: No .env file found, using default environment variables.")
	}

	return &AppConfig{
		ServerPort:  getEnv("SERVER_PORT"),
		DatabaseURI: getEnv("DATABASE_URI"),
	}
}

func getEnv(key string) string {

	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is not set.", key)
	}
	return value
}
