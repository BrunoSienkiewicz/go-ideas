package server

import (
	"encoding/json"
	"fmt"
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

func (s *APIServer) Start() error {
	router := mux.NewRouter()

	router.HandleFunc("/idea", makeHTTPHandleFunc(s.handleIdea))
	router.HandleFunc("/idea/{id}", makeHTTPHandleFunc(s.handleGetIdea))

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

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, APIError{Error: err.Error()})
		}
	}
}

func (s *APIServer) handleAddIdea(w http.ResponseWriter, r *http.Request) error {
	req := new(CreateIdeaRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	idea := NewIdea(req.Name, req.Category, req.Attributes)
	if err := s.store.AddIdea(idea); err != nil {
		return err
	}

	return writeJSON(w, http.StatusCreated, idea)
}

func (s *APIServer) handleGetIdea(w http.ResponseWriter, r *http.Request) error {
	attributes := []Attribute{
		Attribute{Name: "cocktail", Value: "link"},
		Attribute{Name: "plan", Value: "1. spacerek 2. przygotowanie przepisów 3. ubranie się na różowo 4. film"},
	}

	idea := NewIdea("Barbie movie night", "Date", attributes)

	return writeJSON(w, http.StatusOK, idea)
}

func (s *APIServer) handleDeleteIdea(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func writeJSON(w http.ResponseWriter, code int, v any) error {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type APIError struct {
	Error string `json:"error"`
}
