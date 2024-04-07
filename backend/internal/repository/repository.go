package repository

import (
	types "github.com/BrunoSienkiewicz/go_ideas/types"
)

type Repository[T types.DbObject, U types.ApiObject] interface {
	convertToDbObject(obj *U) T
	convertFromDbObject(dbObj *T) U
}

type RepositoryError interface {
	Error() string
}

type NotFoundError struct {
	Err string
}

func (e NotFoundError) Error() string {
	return e.Err
}

type AlreadyExistsError struct {
	Err string
}

func (e AlreadyExistsError) Error() string {
	return e.Err
}
