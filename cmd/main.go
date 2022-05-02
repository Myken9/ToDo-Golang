package main

import (
	todo "github.com/Myken9/ToDo-Golang.git"
	"log"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatal(err)
	}
}
