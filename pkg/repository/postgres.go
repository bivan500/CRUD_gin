package repository

import (
	"github.com/jmoiron/sqlx"
)

const (
	usersTable           = "users"
	usersReadListTable   = "users_list"
	readListsTable       = "lists"
	booksInReadListTable = "books_in_list"
	bookTable            = "book"
)

func NewPostgresDB(DatabaseURL string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", DatabaseURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
