package main

import (
	crudApp "CRUD_GIN"
	"CRUD_GIN/pkg/handler"
	"CRUD_GIN/pkg/repository"
	"CRUD_GIN/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(crudApp.Server)
	err := srv.Run("8080", handlers.InitRoutes())
	if err != nil {
		log.Fatal("error in server start: %v", err.Error())
	}
}
