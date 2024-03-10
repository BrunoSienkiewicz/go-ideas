package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func writeJSON(w http.ResponseWriter, code int, v any) error {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type APIError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Start() error {
	router := mux.NewRouter()

	router.HandleFunc("/idea", makeHTTPHandleFunc(s.handleIdea))

	log.Printf("API server listening on %s", s.listenAddr)

	return http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleIdea(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetIdea(w, r)
	} else if r.Method == "POST" {
		return s.handleAddIdea(w, r)
	} else if r.Method == "DELETE" {
		return s.handleDeleteIdea(w, r)
	}

	return fmt.Errorf("unsupported method %s", r.Method)
}

func (s *APIServer) handleAddIdea(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleGetIdea(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteIdea(w http.ResponseWriter, r *http.Request) error {
	return nil
}
