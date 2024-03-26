package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	storage "github.com/BrunoSienkiewicz/go_ideas/internal/storage"
	types "github.com/BrunoSienkiewicz/go_ideas/types"
)

type IdeaHandler struct {
	store storage.Storage[types.Idea]
}

func (h *IdeaHandler) handleIdea(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return h.handleGetIdea(w, r)
	} else if r.Method == "POST" {
		return h.handleAddIdea(w, r)
	} else if r.Method == "DELETE" {
		return h.handleDeleteIdea(w, r)
	}

	return fmt.Errorf("unsupported method %s", r.Method)
}

func (h *IdeaHandler) handleAddIdea(w http.ResponseWriter, r *http.Request) error {
	req := new(types.CreateIdeaRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		fmt.Println("Error decoding request: ", err)
		return writeJSON(w, http.StatusBadRequest, err)
	}

	idea := types.NewIdea(req.Name, req.Category, req.Attributes)
	if err := h.store.AddObject(idea); err != nil {
		return writeJSON(w, http.StatusInternalServerError, err)
	}

	return writeJSON(w, http.StatusCreated, idea)
}

func (h *IdeaHandler) handleGetIdea(w http.ResponseWriter, r *http.Request) error {
	req := new(types.GetIdeaRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		fmt.Println("Error decoding request: ", err)
		return writeJSON(w, http.StatusBadRequest, err)
	}

	idea, err := h.store.GetObject(req.ID)
	if err != nil {
		return writeJSON(w, http.StatusNotFound, err)
	}

	return writeJSON(w, http.StatusOK, idea)
}

func (h *IdeaHandler) handleDeleteIdea(w http.ResponseWriter, r *http.Request) error {
	return nil
}
