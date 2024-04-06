package api

import (
	"net/http"

	db "github.com/BrunoSienkiewicz/go_ideas/internal/db"
	repository "github.com/BrunoSienkiewicz/go_ideas/internal/repository"
)

func NewRouter(store *db.Postgres) http.Handler {
	r := http.NewServeMux()

	ideaRepository := repository.NewIdeaRepository(store)

	ideaHandler := NewIdeaHandler(ideaRepository)

	r.HandleFunc("/idea", makeHTTPHandleFunc(ideaHandler.handleIdea))
	r.HandleFunc("/idea/{id}", makeHTTPHandleFunc(ideaHandler.handleGetIdeaByID))

	return r
}
