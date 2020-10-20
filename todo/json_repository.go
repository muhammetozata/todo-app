package todo

import (
	"context"

	"encoding/json"
	"net/http"

	"github.com/muhammetozata/todo-app/models"
)

// TJSONRepository ...
type TJSONRepository struct {
	URL string
}

// NewTJSONRepository a construct
// https://jsonplaceholder.typicode.com/todos
func NewTJSONRepository(url string) models.TodoRepository {
	return &TJSONRepository{URL: url}
}

// GetByID ...
func (tj *TJSONRepository) GetByID(ctx context.Context, id int64) (models.Todo, error) {

	var resp *http.Response

	resp, _ = http.Get(tj.URL)

	// defer resp.Body.Close()

	var todo models.Todo

	decoder := json.NewDecoder(resp.Body)

	decoder.Decode(&todo)

	return todo, nil
}
