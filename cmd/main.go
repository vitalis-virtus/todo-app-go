package main

import (
	"log"

	"github.com/vitalis-virtus/todo-app-go"
	"github.com/vitalis-virtus/todo-app-go/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal("error occuredwhile running http server: ", err.Error())
	}
}
