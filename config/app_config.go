package config

import (
	"log"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	HOST string
	Port string
}

var App *AppConfig

func LoadAppConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("No .env file found")
	}

	App = &AppConfig{
		HOST: getEnv("APP_HOST", "localhost"),
		Port: getEnv("APP_PORT", "8000"),
	}
}
