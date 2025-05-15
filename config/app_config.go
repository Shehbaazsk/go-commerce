package config

type AppConfig struct {
	AppHost                string
	AppPort                string
	JWTSecret              string
	JWT_EXPIRATION_MINUTES int
}

var App *AppConfig

func LoadAppConfig() {
	App = &AppConfig{
		AppHost:                getEnv("APP_HOST", "localhost"),
		AppPort:                getEnv("APP_PORT", "8080"),
		JWTSecret:              getEnv("JWT_SECRET", "your_secret_key"),
		JWT_EXPIRATION_MINUTES: getEnvAsInt("JWT_EXPIRATION_MINUTES", 60),
	}
}
