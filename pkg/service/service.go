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
	Create(userId int, list crudApp.ReadList) (int, error)
	GetLists(userId int) ([]crudApp.ReadList, error)
	UpdateList(userId int, listId int, list crudApp.ReadList) (bool, error)
	GetListsById(listId int) (crudApp.ReadList, error)
	DeleteList(userId int, ListId int) (error)
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
		Auth:     NewAuthService(repos.Auth),
		BookList: NewReadListService(repos.BookList),
	}
}
