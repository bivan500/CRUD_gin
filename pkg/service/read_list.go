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

func (s *ReadListService) UpdateList(userId int, listId int, list crudApp.ReadList) (bool, error) {
	return s.repo.UpdateList(userId, listId, list)
}

func (s *ReadListService) GetListsById(listId int) (crudApp.ReadList, error) {
	return s.repo.GetListsById(listId)
}

func (s *ReadListService) DeleteList(userId int, ListId int) error {
	return s.repo.DeleteList(userId, ListId)
}
