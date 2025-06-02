package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	HOST string
	Port string
}

var App *AppConfig

func LoadAppConfig() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get executable path: %v", err)
	}

	// Get the root directory (two levels up inside /cmd)
	projectRoot := filepath.Dir(filepath.Dir(exePath))

	envPath := filepath.Join(projectRoot, ".env")
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading .env file from %s: %v", envPath, err)
	}

	App = &AppConfig{
		HOST: getEnv("APP_HOST", "localhost"),
		Port: getEnv("APP_PORT", "8000"),
	}
}
