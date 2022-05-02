package main

import (
	todo "github.com/Myken9/ToDo-Golang.git"
	"github.com/Myken9/ToDo-Golang.git/pkg/handler"
	"github.com/Myken9/ToDo-Golang.git/pkg/repository"
	"github.com/Myken9/ToDo-Golang.git/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
