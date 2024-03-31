package api

import (
	"net/http"

	"github.com/BrunoSienkiewicz/go_ideas/internal/storage"
)

func NewRouter(store *storage.PostgresStorage) http.Handler {
	r := http.NewServeMux()

	ideaHandler := NewIdeaHandler(store)

	r.HandleFunc("/idea", makeHTTPHandleFunc(ideaHandler.handleIdea))
	r.HandleFunc("/idea/{id}", makeHTTPHandleFunc(ideaHandler.handleGetIdeaByID))

	return r
}
