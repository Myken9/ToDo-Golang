package main

import (
	todo "github.com/Myken9/ToDo-Golang.git"
	"github.com/Myken9/ToDo-Golang.git/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
