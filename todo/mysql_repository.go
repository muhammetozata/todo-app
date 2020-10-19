package todo

import (
	"context"
	"database/sql"
	"log"

	_model "github.com/muhammetozata/todo-app/models"
)

type mysqlTodoRepository struct {
	Conn *sql.DB
}

func NewMysqlTodoRepository(conn *sql.DB) _model.TodoRepository {
	return &mysqlTodoRepository{Conn: conn}
}

func (r *mysqlTodoRepository) query(ctx context.Context, query string, args ...interface{}) (result []_model.Todo, err error) {
	rows, err := r.Conn.QueryContext(ctx, query, args)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	defer rows.Close()

	result = make([]_model.Todo, 0)

	for rows.Next() {
		t := _model.Todo{}

		err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Completed,
		)

		if err != nil {
			log.Fatal(err.Error())
			return nil, err
		}

		result = append(result, t)
	}

	return
}

func (r *mysqlTodoRepository) GetByID(ctx context.Context, id int64) (_model.Todo, error) {

	query := `SELECT id, title, completed FROM todos WHERE ID = ?`

	result, err := r.query(ctx, query)

	if err != nil {
		log.Fatal(err.Error())
		return _model.Todo{}, err
	}

	return result[0], nil
}
