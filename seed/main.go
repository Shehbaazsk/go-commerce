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

type CitiesData struct {
	Name string `json:"name"`
}

type StatesData struct {
	Name   string `json:"name"`
	Cities []CitiesData
}

type CountryData struct {
	Name      string       `json:"name"`
	PhoneCode string       `json:"phonecode"`
	States    []StatesData `json:"states"`
}

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading .env file from %v", err)
	}

	ctx := context.Background()
	pg_pool, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer pg_pool.Close()

	file, err := os.ReadFile("india_state_city_data.json")
	if err != nil {
		log.Fatal("Failed to read JSON file:", err)
	}

	var country CountryData
	if err := json.Unmarshal(file, &country); err != nil {
		log.Fatal("Failed to unmarshal JSON:", err)
	}

	tx, err := pg_pool.Begin(ctx)
	if err != nil {
		log.Fatal("Failed to start transaction:", err)
	}
	defer tx.Rollback(ctx)

	var countryID int
	err = tx.QueryRow(ctx,
		`INSERT INTO countries (name, phone_code) VALUES ($1, $2) RETURNING id`,
		country.Name, country.PhoneCode,
	).Scan(&countryID)
	if err != nil {
		log.Fatal("Failed to insert country:", err)
	}

	for _, state := range country.States {
		var stateID int
		err = tx.QueryRow(ctx,
			`INSERT INTO states (name, country_id) VALUES ($1, $2) RETURNING id`,
			state.Name, countryID,
		).Scan(&stateID)
		if err != nil {
			log.Fatal("Failed to insert state:", err)
		}

		for _, city := range state.Cities {
			_, err = tx.Exec(ctx,
				`INSERT INTO cities (name, state_id) VALUES ($1, $2)`,
				city.Name, stateID,
			)
			if err != nil {
				log.Fatal("Failed to insert city:", err)
			}
		}
	}

	if err := tx.Commit(ctx); err != nil {
		log.Fatal("Failed to commit transaction:", err)
	}

	fmt.Println("Database seeded successfully.")
}
