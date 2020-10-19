package models

import "context"

// Todo ...
type Todo struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	UserID    int64  `json:"userId"`
}

// TodoRepository for Todo struct ...
type TodoRepository interface {
	GetById(ctx context.Context, id int64) (Todo, error)
	GetByTitle(ctx context.Context, title string) (Todo, error)
	GetByUserId(ctx context.Context, userID int64) (Todo, error)
	Store(ctx context.Context, t *Todo) error
	Update(ctx context.Context, t *Todo) error
	Delete(ctx context.Context, id int64) error
}
