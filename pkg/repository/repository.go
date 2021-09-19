package repository

import "github.com/jmoiron/sqlx"

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

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
