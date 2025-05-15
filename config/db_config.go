package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
	DBSSLMode string
}

var DBConf *DBConfig

func LoadDBConfig() {
	DBConf = &DBConfig{
		DBHost:    getEnv("DB_HOST", "localhost"),
		DBPort:    getEnv("DB_PORT", "5432"),
		DBUser:    getEnv("DB_USER", "root"),
		DBPass:    getEnv("DB_PASSWORD", "root"),
		DBName:    getEnv("DB_NAME", "test"),
		DBSSLMode: getEnv("DB_SSLMODE", "disable"),
	}
}

var DB *gorm.DB

func ConnectDB() {
	LoadDBConfig()
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		DBConf.DBHost, DBConf.DBPort, DBConf.DBUser, DBConf.DBPass, DBConf.DBName, DBConf.DBSSLMode,
	)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database
	fmt.Println("Database connected")
	RunMigrations()
}
