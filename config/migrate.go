package config

import (
	"log"

	"github.com/shehbaazsk/go-commerce/models"
)

func RunMigrations() {
	err := DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.UserProfile{},
		&models.Address{},
	)
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}

	log.Println("Database migrated successfully")
}
