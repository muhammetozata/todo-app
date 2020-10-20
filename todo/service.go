package todo

import (
	"context"

	_model "github.com/muhammetozata/todo-app/models"
)

// todoService ...
type todoService struct {
	TodoRepo _model.TodoService
}

// NewTodoService is a construct
func NewTodoService(todoRepo _model.TodoRepository) _model.TodoService {
	return &todoService{
		TodoRepo: todoRepo,
	}
}

func (ts *todoService) GetByID(ctx context.Context, id int64) (_model.Todo, error) {

	ctx, cancel := context.WithTimeout(ctx, 2000)

	defer cancel()

	todo, err := ts.TodoRepo.GetByID(ctx, id)

	if err != nil {
		return todo, err
	}

	return todo, nil
}
