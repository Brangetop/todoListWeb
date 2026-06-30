package main

import (
	"fmt"

	"brange.net/todoListWeb/http"
	"brange.net/todoListWeb/todo"
)

func main() {
	todoList := todo.NewList()
	httpHandlers := http.NewHTTPHandlers(todoList)
	httpServer := http.NewHTTPServer(*httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println(err)
	}
}
