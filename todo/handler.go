package todo

import (
	"github.com/labstack/echo"
	"github.com/muhammetozata/todo-app/models"
)

type TodoHandler struct {
	TUsecase models.TodoUsecase
}

func NewTodoHandler(e *echo.Echo, us models.TodoUsecase) {
	todoHandler := &TodoHandler{
		TUsecase: tUsecase
	}

	e.GET("todos/:id", todoHandler.TUsecase.GetById)
}


func (th *TodoHandler) GetByID() {

}

