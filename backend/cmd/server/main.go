package main

import (
	"log"

	config "github.com/BrunoSienkiewicz/go_ideas/config"
	api "github.com/BrunoSienkiewicz/go_ideas/internal/api"
	db "github.com/BrunoSienkiewicz/go_ideas/internal/db"
	server "github.com/BrunoSienkiewicz/go_ideas/internal/server"
)

func main() {
	cfg := config.NewConfig()

	store, err := db.NewPostgres()
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	router := api.NewRouter(store)

	server := server.NewAPIServer(cfg.ListenAddr)
	server.Start(router)
}
