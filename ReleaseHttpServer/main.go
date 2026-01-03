package main

import (
	"ReleaseHttpServer/http"
	"ReleaseHttpServer/todo"
	"fmt"
)

func main() {
	todoList := todo.NewList()
	handlers := http.NewHTTPHandlers(todoList)
	server := http.NewHTTPServer(handlers)

	if err := server.StartServer(); err != nil {
		fmt.Println("Faled to start server", err)
	}
}
