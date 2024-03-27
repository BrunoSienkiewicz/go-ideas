package main

import (
	"log"

	api "github.com/BrunoSienkiewicz/go_ideas/internal/api"
	server "github.com/BrunoSienkiewicz/go_ideas/internal/server"
	storage "github.com/BrunoSienkiewicz/go_ideas/internal/storage"
)

func main() {
	store, err := storage.NewPostgresStorage()
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	router := api.NewRouter()

	server := server.NewAPIServer(":5000", *store)
	server.Start(router)
}
