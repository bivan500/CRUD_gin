package service

import (
	crudApp "CRUD_GIN"
	"CRUD_GIN/pkg/repository"
)

type Auth interface {
	CreateUser(user crudApp.User) (int, error)
	GenerateToken(user crudApp.User) (string, error)
	ParseToken(token string) (int, error)
}

type BookList interface {
}

type Book interface {
}

type Service struct {
	Auth
	BookList
	Book
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repos.Auth),
	}
}
