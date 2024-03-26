package api

import (
	"net/http"
)

func NewRouter() http.Handler {
	r := http.NewServeMux()

	ideaHandler := &IdeaHandler{}

	r.HandleFunc("/idea", makeHTTPHandleFunc(ideaHandler.handleIdea))

	return r
}
