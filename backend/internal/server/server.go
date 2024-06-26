package server

import (
	"log"
	"net/http"
)

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Start(router http.Handler) error {
	log.Printf("API server listening on %s", s.listenAddr)

	return http.ListenAndServe(s.listenAddr, router)
}
