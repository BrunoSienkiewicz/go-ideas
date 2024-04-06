package repository

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
