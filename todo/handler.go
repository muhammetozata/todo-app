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

// TodooHandler ...
type TodooHandler struct {
	TService _model.TodoService
}

// NewTodooHandler is a construct
func NewTodooHandler(e *echo.Echo, ts _model.TodoService) {
	todooHandler := &TodooHandler{
		TService: ts,
	}

	e.GET("/todos/:id", todooHandler.GetByID)
}

// GetByID ...
func (th *TodooHandler) GetByID(c echo.Context) error {
	idStr := c.Param("id")

	id, _ := strconv.ParseInt(idStr, 10, 64)

	ctx := c.Request().Context()

	todo, err := th.TService.GetByID(ctx, id)

	if err != nil {
		return c.JSON(http.StatusNotFound, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, todo)
}
