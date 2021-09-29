package main

import (
	crudApp "CRUD_GIN"
	"CRUD_GIN/pkg/handler"
	"CRUD_GIN/pkg/repository"
	"CRUD_GIN/pkg/service"
	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	db, err := repository.NewPostgresDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		logrus.Fatal("error in init db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(crudApp.Server)
	err = srv.Run(os.Getenv("PORT"), handlers.InitRoutes())
	if err != nil {
		logrus.Fatal("error in server start: %v", err.Error())
	}
}
