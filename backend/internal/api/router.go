package api

import (
	"net/http"

	"github.com/BrunoSienkiewicz/go_ideas/internal/storage"
	attributestorage "github.com/BrunoSienkiewicz/go_ideas/internal/storage/attributestorage"
	ideastorage "github.com/BrunoSienkiewicz/go_ideas/internal/storage/ideastorage"
)

func NewRouter(store *storage.PostgresStorage) http.Handler {
	r := http.NewServeMux()

	attributeStorage := attributestorage.NewAttributeStorage(store)
	attributeHandler := NewAttributeHandler(attributeStorage)

	ideaStorage := ideastorage.NewIdeaStorage(store, attributeStorage)
	ideaHandler := NewIdeaHandler(ideaStorage)

	r.HandleFunc("/idea", makeHTTPHandleFunc(ideaHandler.handleIdea))
	r.HandleFunc("/idea/{id}", makeHTTPHandleFunc(ideaHandler.handleGetIdeaByID))

	return r
}
