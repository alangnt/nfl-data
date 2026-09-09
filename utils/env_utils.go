package utils

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func GetSportradarAPIKey() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", errors.New("Error loading .env file")
	}

	sportradarKey := os.Getenv("SPORTRADAR_API_KEY")
	if sportradarKey == "" {
		return "", errors.New("API key required")
	}

	return sportradarKey, nil
}
