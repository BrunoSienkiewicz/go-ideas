package main

import (
	"log"
)

func main() {
	store, err := NewPostgresStorage()
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	server := NewAPIServer(":3000", store)
	server.Start()
}
