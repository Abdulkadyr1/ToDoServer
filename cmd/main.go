package main

import (
	"ToDo/internal/config"
	tasks2 "ToDo/tasks"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// INSERT INTO products (name, price) VALUES ('Sample Product', 19.99);
func main() {
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	err = db.AutoMigrate(tasks2.Task{})
	if err != nil {
		log.Fatal("Failed to migrate database")
	}
	config.DB = db
	e := echo.New()
	TaskHandler := &tasks2.TasksHandler{}
	e.GET("/tasks", TaskHandler.GetAllTasks)
	e.POST("/tasks", TaskHandler.PostTask)
	e.PUT("/tasks/:id", TaskHandler.UpdateTask)
	e.DELETE("tasks/:id", TaskHandler.DeleteTask)
	err = e.Start(":8080")
	if err != nil {
		log.Fatal("Error starting server:")
	}
}
