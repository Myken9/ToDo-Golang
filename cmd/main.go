package main

import (
	todo "github.com/Myken9/ToDo-Golang.git"
	"github.com/Myken9/ToDo-Golang.git/pkg/handler"
	"github.com/Myken9/ToDo-Golang.git/pkg/repository"
	"github.com/Myken9/ToDo-Golang.git/pkg/service"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Post:     "5432",
		Username: "valery",
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   "tododb",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to initiize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
