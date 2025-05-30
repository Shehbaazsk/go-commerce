package config

import "os"

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
