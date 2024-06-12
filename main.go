package main

import (
	"ToDo/handlers"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	e := echo.New()
	e.GET("/tasks", handlers.GetTasks)
	err := e.Start(":8080")
	if err != nil {
		log.Fatal("Error starting server:")
	}
}
