package attributestorage

import (
	_ "github.com/lib/pq"

	storage "github.com/BrunoSienkiewicz/go_ideas/internal/storage"
)

type AttributeStorage struct {
	store *storage.PostgresStorage
}

func NewAttributeStorage(postgres *storage.PostgresStorage) *AttributeStorage {
	return &AttributeStorage{
		store: postgres,
	}
}
