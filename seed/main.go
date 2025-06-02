package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shehbaazsk/go-commerce/config"
)

// Data models the outer structure
type Data map[string]CountryDetail

type CountryDetail struct {
	PhoneCode string                `json:"phonecode"`
	States    []map[string][]string `json:"State"`
}

func main() {
	// Load environment variables
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connect to PostgreSQL DB
	ctx := context.Background()
	pgPool, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer pgPool.Close()

	// Read the JSON file
	file, err := os.ReadFile("india_state_city_data.json")
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	// Parse JSON
	var data Data
	if err := json.Unmarshal(file, &data); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Start DB transaction
	tx, err := pgPool.Begin(ctx)
	if err != nil {
		log.Fatalf("Failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	// Process each country (in your example, just "India")
	for countryName, country := range data {
		var countryID int
		// Insert or get existing country
		err := tx.QueryRow(ctx,
			`INSERT INTO countries (name, phone_code)
			 VALUES ($1, $2)
			 ON CONFLICT (name) DO UPDATE SET phone_code = EXCLUDED.phone_code
			 RETURNING id`,
			countryName, country.PhoneCode,
		).Scan(&countryID)
		if err != nil {
			log.Fatalf("Failed to insert/get country '%s': %v", countryName, err)
		}
		fmt.Printf("Country: %s (ID: %d)\n", countryName, countryID)

		for _, stateGroup := range country.States {
			for stateName, cities := range stateGroup {
				var stateID int
				// Insert or get existing state
				err := tx.QueryRow(ctx,
					`INSERT INTO states (name, country_id)
					 VALUES ($1, $2)
					 ON CONFLICT (name, country_id) DO NOTHING
					 RETURNING id`,
					stateName, countryID,
				).Scan(&stateID)

				// If state already exists, fetch its ID
				if err != nil {
					err = tx.QueryRow(ctx,
						`SELECT id FROM states WHERE name=$1 AND country_id=$2`,
						stateName, countryID,
					).Scan(&stateID)
					if err != nil {
						log.Fatalf("Failed to get existing state '%s': %v", stateName, err)
					}
				}
				fmt.Printf("  State: %s (ID: %d)\n", stateName, stateID)

				for _, cityName := range cities {
					// Insert or skip existing city
					_, err := tx.Exec(ctx,
						`INSERT INTO cities (name, state_id)
						 VALUES ($1, $2)
						 ON CONFLICT (name, state_id) DO NOTHING`,
						cityName, stateID,
					)
					if err != nil {
						log.Fatalf("Failed to insert city '%s': %v", cityName, err)
					}
					fmt.Printf("    City: %s\n", cityName)
				}
			}
		}
	}

	// Commit transaction
	if err := tx.Commit(ctx); err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}

	fmt.Println("Database seeding completed successfully.")
}
