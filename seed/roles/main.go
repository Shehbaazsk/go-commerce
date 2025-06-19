package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/shehbaazsk/go-commerce/config"
)

type Role struct {
	Name        string
	Description string
}

func main() {
	roles := []Role{
		{"ADMIN", "Admin Level Access"},
		{"STAFF", "Staff Level Access"},
		{"SELLER", "Seller Level Access"},
		{"CUSTOMER", "Customer Level Access"},
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file from  %v", err)
	}

	// Connect to PostgreSQL DB
	ctx := context.Background()
	pgPool, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer pgPool.Close()
	log.Println("Connected to the database")
	tx, err := pgPool.Begin(ctx)
	defer tx.Rollback(ctx)

	if err != nil {
		log.Fatalf("Failed to start transaction: %v", err)
	}
	for _, role := range roles {
		tag, err := tx.Exec(ctx,
			`INSERT INTO roles (name, description) VALUES ($1, $2)`,
			role.Name, role.Description,
		)
		fmt.Println(tag)
		if err != nil {
			log.Fatalf("Failed to insert role: %v", err)
		}
	}
	if err := tx.Commit(ctx); err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}
	fmt.Println("Roles seeded successfully")

}
