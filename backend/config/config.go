package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadConfig loads environment variables
func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

// GetEnv returns the value of an environment variable
func GetEnv(key string) string {
	return os.Getenv(key)
}
