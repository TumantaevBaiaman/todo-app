package main

import (
	"log"
	todo_app "todo-app"
	"todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo_app.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}
