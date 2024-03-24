package server

import (
	"internal/api"
	"internal/db"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Start(router *mux.Router) error {
	log.Printf("API server listening on %s", s.listenAddr)

	return http.ListenAndServe(s.listenAddr, router)
}
