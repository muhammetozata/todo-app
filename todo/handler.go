package todo

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	_model "github.com/muhammetozata/todo-app/models"
)

// ResponseError ...
type ResponseError struct {
	Message string `json:"message"`
}

// TodoHandler ...
type TodoHandler struct {
	TService _model.TodoService
}

func NewTodoHandler(e *echo.Echo, ts _model.TodoService) {
	todoHandler := &TodoHandler{
		TService: ts,
	}

	e.GET("/todos/:id", todoHandler.GetByID)
}

func (th *TodoHandler) GetByID(c echo.Context) error {
	// idP, err := strconv.Atoi(c.Param("id"))

	idP, err := strconv.Atoi(c.Param("id"))

	if err != nil {

		return c.JSON(http.StatusNotFound, ResponseError{Message: err.Error()})
	}

	id := int64(idP)

	ctx := c.Request().Context()

	todo, err := th.TService.GetByID(ctx, id)

	if err != nil {
		return c.JSON(http.StatusNotFound, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, todo)
}
