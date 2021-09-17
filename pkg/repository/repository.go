package repository

type Auth interface {
}

type BookList interface {
}

type Book interface {
}

type Repository struct {
	Auth
	BookList
	Book
}

func NewRepository() *Repository {
	return &Repository{}
}
