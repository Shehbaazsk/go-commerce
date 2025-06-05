package config

import (
	"log"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	HOST                   string
	Port                   string
	JWT_SECRET             string
	JWT_EXPIRATION         int
	JWT_REFRESH_EXPIRATION int
}

var App *AppConfig

func LoadAppConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file from : %v", err)
	}

	App = &AppConfig{
		HOST:                   getEnv("APP_HOST", "localhost"),
		Port:                   getEnv("APP_PORT", "8000"),
		JWT_SECRET:             getEnv("JWT_SECRET", ""),
		JWT_EXPIRATION:         getEnvInt("JWT_EXPIRATION", 30),
		JWT_REFRESH_EXPIRATION: getEnvInt("JWT_REFRESH_EXPIRATION", 1440),
	}
}
