package main

import (
	"ToDo/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/tasks", handlers.GetTasks)
	e.Start(":8080")
}
