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

func NewIdeaHandler(postgres *storage.PostgresStorage) *IdeaHandler {
	idea_storage := storage.NewIdeaStorage(postgres)

	return &IdeaHandler{
		store: idea_storage,
	}
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
	ideas, err := h.store.ListObjects()
	if err != nil {
		return writeJSON(w, http.StatusInternalServerError, err)
	}

	return writeJSON(w, http.StatusOK, ideas)
}

func (h *IdeaHandler) handleDeleteIdea(w http.ResponseWriter, r *http.Request) error {
	req := new(types.DeleteIdeaRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return writeJSON(w, http.StatusBadRequest, err)
	}

	if err := h.store.DeleteObject(req.ID); err != nil {
		return writeJSON(w, http.StatusInternalServerError, err)
	}

	return writeJSON(w, http.StatusNoContent, nil)
}

func (h *IdeaHandler) handleUpdateIdea(w http.ResponseWriter, r *http.Request) error {
	req := new(types.UpdateIdeaRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return writeJSON(w, http.StatusBadRequest, err)
	}

	idea := types.NewIdea(req.Name, req.Category, req.Attributes)
	idea.ID = req.ID
	if err := h.store.UpdateObject(idea); err != nil {
		return writeJSON(w, http.StatusInternalServerError, err)
	}

	return writeJSON(w, http.StatusOK, idea)
}

func (h *IdeaHandler) handleGetIdeaByID(w http.ResponseWriter, r *http.Request) error {
	id, err := getIDFromURL(r)
	if err != nil {
		return writeJSON(w, http.StatusBadRequest, err)
	}

	idea, err := h.store.GetObject(id)
	if err != nil {
		return writeJSON(w, http.StatusInternalServerError, err)
	}

	return writeJSON(w, http.StatusOK, idea)
}
