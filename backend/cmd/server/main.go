package main

import (
	"log"

	config "github.com/BrunoSienkiewicz/go_ideas/config"
	api "github.com/BrunoSienkiewicz/go_ideas/internal/api"
	server "github.com/BrunoSienkiewicz/go_ideas/internal/server"
	storage "github.com/BrunoSienkiewicz/go_ideas/internal/storage"
)

func main() {
	cfg := config.NewConfig()

	store, err := storage.NewPostgresStorage()
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	router := api.NewRouter(store)

	server := server.NewAPIServer(cfg.ListenAddr, *store)
	server.Start(router)
}
