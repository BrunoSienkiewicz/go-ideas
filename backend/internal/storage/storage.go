package storage

import (
	types "github.com/BrunoSienkiewicz/go_ideas/types"
)

type Storage[T types.DbObject] interface {
	Getter[T]
	Adder[T]
	Updater[T]
	Deleter[T]
}

type Getter[T types.DbObject] interface {
	GetObject(id int) (*T, error)
	GetAllObjects() ([]*T, error)
}

type Adder[T types.DbObject] interface {
	AddObject(obj *T) (*T, error)
}

type Updater[T types.DbObject] interface {
	UpdateObject(obj *T) (*T, error)
	UpdateObjectField(id int, field string, value string) error
}

type Deleter[T types.DbObject] interface {
	DeleteObject(id int) error
}
