package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetSportraderAPIKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sportradarKey := os.Getenv("SPORTRADAR_API_KEY")
	if sportradarKey == "" {
		log.Fatal("API key required")
	}

	return sportradarKey
}
