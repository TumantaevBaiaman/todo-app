package main

import (
	"log"
	todo_app "todo-app"
)

func main() {
	srv := new(todo_app.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}
