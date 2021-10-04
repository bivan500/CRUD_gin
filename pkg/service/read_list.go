package service

import (
	crudApp "CRUD_GIN"
	"CRUD_GIN/pkg/repository"
)

type ReadListService struct {
	repo repository.BookList
}

func NewReadListService(repo repository.BookList) *ReadListService {
	return &ReadListService{repo: repo}
}

func (s *ReadListService) Create(userId int, list crudApp.ReadList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *ReadListService) GetLists(userId int) ([]crudApp.ReadList, error) {
	return s.repo.GetLists(userId)
}
