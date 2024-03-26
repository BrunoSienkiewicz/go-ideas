package migration

import (
	"fmt"
	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"os"
	"path/filepath"

	config "github.com/BrunoSienkiewicz/go_ideas/config"
)

func main() {
	// Load the environment variables
	cfg := config.NewConfig()
	connString := cfg.GetDbConnectionString()

	// Create a new instance of the migration tool
	m, err := migrate.New(
		"file://migrations", // Path to the migrations directory
		connString,
	)
	if err != nil {
		log.Fatalf("Error creating migration instance: %v", err)
	}

	// Run the migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error applying migrations: %v", err)
	}

	fmt.Println("Migrations applied successfully")
}
