package main

import (
	crudApp "CRUD_GIN"
	"CRUD_GIN/pkg/handler"
	"CRUD_GIN/pkg/repository"
	"CRUD_GIN/pkg/service"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatal("error in init config: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("error in load env file: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.name"),
		SSLMode:  viper.GetString("db.SSLMode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatal("error in init db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(crudApp.Server)
	err = srv.Run(viper.GetString("port"), handlers.InitRoutes())
	if err != nil {
		logrus.Fatal("error in server start: %v", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
