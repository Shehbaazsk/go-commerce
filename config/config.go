package config

import (
	"os"
	"strconv"
)

func InitConfig() {
	LoadAppConfig()
	LoadDBConfig()
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

func getEnvInt(key string, fallback int) int {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	intValue, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return intValue
}
