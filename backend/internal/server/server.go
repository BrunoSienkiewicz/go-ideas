package server

import (
	"log"
	"net/http"

	storage "github.com/BrunoSienkiewicz/go_ideas/internal/storage"
)

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string, store storage.PostgresStorage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Start(router http.Handler) error {
	log.Printf("API server listening on %s", s.listenAddr)

	return http.ListenAndServe(s.listenAddr, router)
}
