package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Error loading .env file, using environment variables instead.")
	}

	LoadAppConfig()
	LoadDBConfig()
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
