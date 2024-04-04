package repository

import (
	storage "github.com/BrunoSienkiewicz/go_ideas/internal/storage"
	types "github.com/BrunoSienkiewicz/go_ideas/types"
)

type IdeaRepository struct {
	categoryStorage  storage.Storage[types.DbCategory]
	attributeStorage storage.Storage[types.DbAttribute]
	ideaStorage      storage.Storage[types.DbIdea]
}
