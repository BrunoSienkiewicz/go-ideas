package api

import (
	"encoding/json"
	"fmt"
	"internal/storage"
	"net/http"
)

type Idea struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Category   string      `json:"category"`
	Attributes []Attribute `json:"attributes"`
}

type Attribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CreateIdeaRequest struct {
	Name       string      `json:"name"`
	Category   string      `json:"category"`
	Attributes []Attribute `json:"attributes"`
}

func NewIdea(name string, category string, attributes []Attribute) *Idea {
	return &Idea{
		Name:       name,
		Category:   category,
		Attributes: attributes,
	}
}

type IdeaController struct {
	storage *IdeaStorage
}

func NewIdeaController(storage *IdeaStorage) *IdeaController {
	return &IdeaController{storage: storage}
}

func (c *IdeaController) handleIdea(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return c.handleGetIdea(w, r)
	} else if r.Method == "POST" {
		return c.handleAddIdea(w, r)
	} else if r.Method == "DELETE" {
		return c.handleDeleteIdea(w, r)
	}

	return fmt.Errorf("unsupported method %s", r.Method)
}

func (c *IdeaController) handleAddIdea(w http.ResponseWriter, r *http.Request) error {
	req := new(CreateIdeaRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	idea := NewIdea(req.Name, req.Category, req.Attributes)

	return writeJSON(w, http.StatusCreated, idea)
}

func (c *IdeaController) handleGetIdea(w http.ResponseWriter, r *http.Request) error {
	attributes := []Attribute{
		Attribute{Name: "cocktail", Value: "link"},
		Attribute{Name: "plan", Value: "1. spacerek 2. przygotowanie przepisów 3. ubranie się na różowo 4. film"},
	}

	idea := NewIdea("Barbie movie night", "Date", attributes)

	return writeJSON(w, http.StatusOK, idea)
}

func (c *IdeaController) handleDeleteIdea(w http.ResponseWriter, r *http.Request) error {
	return nil
}