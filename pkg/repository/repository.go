package repository

import (
	crudApp "CRUD_GIN"

	"github.com/jmoiron/sqlx"
)

type Auth interface {
	CreateUser(user crudApp.User) (int, error)
	GetUser(username string, password string) (crudApp.User, error)
}

type BookList interface {
	Create(userId int, list crudApp.ReadList) (int, error)
	GetLists(userId int) ([]crudApp.ReadList, error)
}

type Book interface {
}

type Repository struct {
	Auth
	BookList
	Book
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth:     newAuthPostgres(db),
		BookList: NewListPostgres(db),
	}
}
