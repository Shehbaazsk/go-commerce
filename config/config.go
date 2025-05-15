package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()
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

func getEnvAsInt(key string, fallback int) int {
	if valueStr, exists := os.LookupEnv(key); exists {
		var value int
		_, err := fmt.Sscanf(valueStr, "%d", &value)
		if err == nil {
			return value
		}
	}
	return fallback
}
