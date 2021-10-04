package crudApp

type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

type ReadList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"requierd"`
	Description string `json:"description" db:"description"`
}

type UsersReadList struct {
	Id         int
	UserId     int
	ReadListId int
}
