package main

import (
	crudApp "CRUD_GIN"
	"CRUD_GIN/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(crudApp.Server)
	err := srv.Run("8080", handlers.InitRoutes())
	if err != nil {
		log.Fatal("error in server start: %v", err.Error())
	}
}
