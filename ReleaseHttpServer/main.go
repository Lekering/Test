package main

import (
	"ReleaseHttpServer/http"
	simpleconnect "ReleaseHttpServer/simple_connect"
	simpletable "ReleaseHttpServer/simple_table"
	"ReleaseHttpServer/todo"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	conn, err := simpleconnect.Connect(ctx)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	if err := simpletable.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	todoList := todo.NewList(ctx, conn)
	handlers := http.NewHTTPHandlers(todoList)
	server := http.NewHTTPServer(handlers)

	if err := server.StartServer(); err != nil {
		fmt.Println("Faled to start server", err)
	}
}
