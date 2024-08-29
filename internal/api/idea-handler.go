package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	repository "github.com/BrunoSienkiewicz/go_ideas/internal/repository"
	types "github.com/BrunoSienkiewicz/go_ideas/types"
)

type IdeaHandler struct {
	repository repository.Repository[types.DbIdea, types.Idea]
}

func NewIdeaHandler(repository *repository.IdeaRepository) *IdeaHandler {
	return &IdeaHandler{
		repository: repository,
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

	_attrubutes := make([]types.Attribute, len(req.Attributes))
	for i, attr := range req.Attributes {
		_attrubutes[i] = *types.NewAttribute(attr.Name, attr.Value)
	}

	// TODO: handle idea from request
	_idea := types.NewIdea(req.Name, req.Category, _attrubutes)
	err, idea := h.repository.AddObject(_idea)
	if err != nil {
		return writeJSON(w, http.StatusInternalServerError, err)
	}

	return writeJSON(w, http.StatusCreated, idea)
}

func (h *IdeaHandler) handleGetIdea(w http.ResponseWriter, r *http.Request) error {
	ideas, err := h.repository.GetAllObjects()
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

	if err := h.repository.DeleteObject(req.ID); err != nil {
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
	_idea, err := h.repository.UpdateObject(idea)
	if err != nil {
		return writeJSON(w, http.StatusInternalServerError, err)
	}

	return writeJSON(w, http.StatusOK, _idea)
}

func (h *IdeaHandler) handleGetIdeaByID(w http.ResponseWriter, r *http.Request) error {
	id, err := getIDFromURL(r)
	if err != nil {
		return writeJSON(w, http.StatusBadRequest, err)
	}

	idea, err := h.repository.GetObject(id)
	if err != nil {
		return writeJSON(w, http.StatusInternalServerError, err)
	}

	return writeJSON(w, http.StatusOK, idea)
}
