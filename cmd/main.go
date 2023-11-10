package main

import (
	"log"
	"todolist"
	"todolist/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todolist.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while runnin https server: %s", err.Error())
	}
}
