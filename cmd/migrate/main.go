package main

import (
	"flag"
	"fmt"

	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"

	config "github.com/BrunoSienkiewicz/go_ideas/config"
)

func main() {
	// Parse the command line arguments
	direction := flag.String("direction", "up", "The migration direction (up or down)")
	flag.Parse()

	// Load the environment variables
	cfg := config.NewConfig()
	connString := cfg.GetDbConnectionString()

	// Create a new instance of the migration tool
	m, err := migrate.New(
		"file:///migrations", // Path to the migrations directory
		connString,
	)
	if err != nil {
		log.Fatalf("Error creating migration instance: %v", err)
	}

	var directionFunc func() error
	switch *direction {
	case "up":
		directionFunc = m.Up
	case "down":
		directionFunc = m.Down
	default:
		log.Fatalf("Invalid migration direction: %s (valid values: 'up' or 'down')", *direction)
	}

	// Run the migrations
	if err := directionFunc(); err != nil {
		log.Fatalf("Error applying migrations: %v", err)
	}

	fmt.Println("Migrations applied successfully")
}
