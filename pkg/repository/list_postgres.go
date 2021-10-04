package repository

import (
	crudApp "CRUD_GIN"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ListPostgres struct {
	db *sqlx.DB
}

func NewListPostgres(db *sqlx.DB) *ListPostgres {
	return &ListPostgres{db: db}
}

func (r *ListPostgres) Create(userId int, list crudApp.ReadList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createReadListQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", readListsTable)
	row := r.db.QueryRow(createReadListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersReadListQuery := fmt.Sprintf("INSERT INTO %s (users_id, read_list_id) values ($1, $2) RETURNING id", usersReadListTable)
	_, err = tx.Exec(createUsersReadListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *ListPostgres) GetLists(userId int) ([]crudApp.ReadList, error) {
	var result []crudApp.ReadList
	getReadListQuery := fmt.Sprintf("SELECT lists.id,title,description FROM %s lists INNER JOIN %s user_lists on lists.id = user_lists.list_id WHERE user_lists.users_id = $1", readListsTable, usersReadListTable)
	//getReadListQuery := "SELECT lists.id,title,description FROM lists INNER JOIN user_lists on lists.id = user_lists.list_id WHERE user_lists.users_id = 1"
	err := r.db.Select(&result, getReadListQuery, userId)

	return result, err
}
