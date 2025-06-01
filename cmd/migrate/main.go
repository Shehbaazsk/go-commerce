package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/joho/godotenv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Please provide a command: create, up, down, steps, version, force")
	}
	command := os.Args[1]

	// Load .env file
	if err := godotenv.Load("./.env"); err != nil {
		log.Println("Warning: .env file not found, reading from environment")
	}

	dbURL := getEnv("DATABASE_URL", "")
	if dbURL == "" {
		log.Fatal("DATABASE_URL must be set in .env or environment variables")
	}
	dir, err := filepath.Abs("db/migrations")
	if err != nil {
		log.Fatal(err)
	}
	migrationsPath := "file://" + dir

	switch command {
	case "create":
		if len(os.Args) < 3 {
			log.Fatal("Usage: migrate create <name>")
		}
		name := os.Args[2]
		createMigration(migrationsPath, name)

	case "up":
		migrateUp(migrationsPath, dbURL)

	case "down":
		migrateDown(migrationsPath, dbURL)

	case "steps":
		if len(os.Args) < 3 {
			log.Fatal("Usage: migrate steps <n>")
		}
		n, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Invalid number of steps: %v", err)
		}
		migrateSteps(migrationsPath, dbURL, n)

	case "version":
		showVersion(migrationsPath, dbURL)

	case "force":
		if len(os.Args) < 3 {
			log.Fatal("Usage: migrate force <version>")
		}
		version, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Invalid version number: %v", err)
		}
		forceVersion(migrationsPath, dbURL, version)

	default:
		log.Fatalf("Unknown command: %s", command)
	}
}

func createMigration(migrationsPath, name string) {
	// Requires migrate CLI installed in PATH
	name = strings.ReplaceAll(name, " ", "_")
	cmd := exec.Command("migrate", "create", "-ext", "sql", "-dir", migrationsPath, "-seq", name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to create migration: %v", err)
	}
	fmt.Println("Migration created successfully.")
}

func migrateUp(migrationsPath, dbURL string) {
	m, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		log.Fatalf("Failed to initialize migrate: %v", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration up failed: %v", err)
	}
	fmt.Println("Migrations applied successfully.")
}

func migrateDown(migrationsPath, dbURL string) {
	m, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		log.Fatalf("Failed to initialize migrate: %v", err)
	}
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration down failed: %v", err)
	}
	fmt.Println("Migrations rolled back successfully.")
}

func migrateSteps(migrationsPath, dbURL string, steps int) {
	m, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		log.Fatalf("Failed to initialize migrate: %v", err)
	}
	if err := m.Steps(steps); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration steps failed: %v", err)
	}
	fmt.Printf("Migration steps %d applied successfully.\n", steps)
}

func showVersion(migrationsPath, dbURL string) {
	m, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		log.Fatalf("Failed to initialize migrate: %v", err)
	}
	version, dirty, err := m.Version()
	if err != nil {
		if err == migrate.ErrNilVersion {
			fmt.Println("No migrations applied yet.")
			return
		}
		log.Fatalf("Failed to get migration version: %v", err)
	}
	fmt.Printf("Current migration version: %d, dirty: %v\n", version, dirty)
}

func forceVersion(migrationsPath, dbURL string, version int) {
	m, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		log.Fatalf("Failed to initialize migrate: %v", err)
	}
	if err := m.Force(version); err != nil {
		log.Fatalf("Failed to force migration version: %v", err)
	}
	fmt.Printf("Forced migration version to %d successfully.\n", version)
}
