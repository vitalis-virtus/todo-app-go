package main

import (
	"log"

	"github.com/vitalis-virtus/todo-app-go"
	"github.com/vitalis-virtus/todo-app-go/pkg/handler"
	"github.com/vitalis-virtus/todo-app-go/pkg/repository"
	"github.com/vitalis-virtus/todo-app-go/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal("error occuredwhile running http server: ", err.Error())
	}
}
