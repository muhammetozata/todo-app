package main

import (
	"github.com/labstack/echo"
	"github.com/muhammetozata/todo-app/todo"
)

func main() {

	e := echo.New()

	// mysqlRepo := todo.NewMysqlTodoRepository(nil)

	jsonRepo := todo.TJSONRespository("https://jsonplaceholder.typicode.com/todos/1")
	todoService := todo.NewTodoService(jsonRepo)

	todo.NewTodooHandler(e, todoService)

	e.Start(":9000")

}
