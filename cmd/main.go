package main

import (
	"log"
	"todolist"
	"todolist/pkg/handler"
	"todolist/pkg/repository "
	"todolist/pkg/service"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(): err ! = nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todolist.Server)
	log.Fatalf("error occured while running https server: %s", srv.Run("8080", handlers.InitRoutes()).Error())

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
