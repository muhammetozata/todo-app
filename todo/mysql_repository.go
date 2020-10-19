package todo

import (
	"database/sql"

	_model "../models"
)

type mysqlTodoRepository struct {
	Conn *sql.DB
}

func NewMysqlTodoRepository(conn *sql.DB) _model.TodoRepository {
	return &mysqlTodoRepository{Conn}
}

func (r *mysqlTodoRepository) GetById() (_model.Todo, error) {

}
