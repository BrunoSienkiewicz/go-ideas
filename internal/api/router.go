package api

import (
	"internal/storage"
	"net/http"
)

func NewRouter() http.Handler {
	r := http.NewServeMux()

	ideaController := NewIdeaController(NewIdeaStorage())

	r.HandleFunc("/idea", makeHTTPHandleFunc(ideaController.handleIdea))

	return r
}
