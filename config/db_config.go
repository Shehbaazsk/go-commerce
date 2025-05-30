package config

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

var DB *DBConfig

func LoadDBConfig() {
	DB = &DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		Username: getEnv("DB_USERNAME", "user"),
		Password: getEnv("DB_PASSWORD", "password"),
		Database: getEnv("DB_DATABASE", "mydb"),
	}
}
