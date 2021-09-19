package main

import (
	crudApp "CRUD_GIN"
	"CRUD_GIN/pkg/handler"
	"CRUD_GIN/pkg/repository"
	"CRUD_GIN/pkg/service"
	"log"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("error in init config: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(crudApp.Server)
	err := srv.Run(viper.GetString("port"), handlers.InitRoutes())
	if err != nil {
		log.Fatal("error in server start: %v", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
