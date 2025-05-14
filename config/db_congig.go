package config

type DBConfig struct {
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
	DBSSLMode string
}

var DB *DBConfig

func LoadDBConfig() {
	DB = &DBConfig{
		DBHost:    getEnv("DB_HOST", "localhost"),
		DBPort:    getEnv("DB_PORT", "5432"),
		DBUser:    getEnv("DB_USER", "root"),
		DBPass:    getEnv("DB_PASSWORD", "root"),
		DBName:    getEnv("DB_NAME", "go_commerce"),
		DBSSLMode: getEnv("DB_SSLMODE", "disable"),
	}
}
