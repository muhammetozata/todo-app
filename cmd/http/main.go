package main

import (
	"net/url"

	"github.com/labstack/echo"
	"github.com/muhammetozata/todo-app/todo"
)

func main() {

	e := echo.New()

	urlP, _ := url.Parse("https://jsonplaceholder.typicode.com")
	urlStr := string(urlP.String())

	jsonRepo := todo.NewTJSONRepository(urlStr)
	todoService := todo.NewTodoService(jsonRepo)

	todo.NewTodooHandler(e, todoService)

	e.Start(":9000")

}
