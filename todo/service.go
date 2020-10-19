package todo

import (
	"context"

	_model "github.com/muhammetozata/todo-app/models"
)

type todoService struct {
	TodoRepo _model.TodoRepository
}

func NewTodoService(todoRepo _model.TodoRepository) _model.TodoService {

	return &todoService{
		TodoRepo: todoRepo
	}
}


func (ts *todoService) GetByID(ctx context.Context, id int64) _model.Todo, error {

	ts.TodoRepo.GetByID(ctx, id)

}