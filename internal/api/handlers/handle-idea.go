package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
	req := new(GetIdeaRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	idea, err := s.store.GetIdea(req.ID)

	return writeJSON(w, http.StatusOK, idea)
}

func (s *APIServer) handleDeleteIdea(w http.ResponseWriter, r *http.Request) error {
	return nil
}
