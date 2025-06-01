package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	HOST string
	Port string
}

var App *AppConfig

func LoadAppConfig() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}

	// Construct full path to .env relative to project root
	envPath := wd + "/.env"

	// Load .env
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading .env file from %s: %v", envPath, err)
	}

	App = &AppConfig{
		HOST: getEnv("APP_HOST", "localhost"),
		Port: getEnv("APP_PORT", "8000"),
	}
}
