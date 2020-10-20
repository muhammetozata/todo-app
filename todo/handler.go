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

	idP := c.Param("id")
	id, _ := strconv.Atoi(idP)

	id64 := int64(id)

	ctx := c.Request().Context()

	todo, _ := th.TService.GetByID(ctx, id64)

	return c.JSON(http.StatusOK, todo)
}
