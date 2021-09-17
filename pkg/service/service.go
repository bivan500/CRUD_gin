package service

import "CRUD_GIN/pkg/repository"

type Auth interface {
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

func NewService(r *repository.Repository) *Service {
	return &Service{}
}
