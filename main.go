package main

import (
	"ToDo/config"
	"ToDo/handlers"
	"ToDo/models"
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
	err = db.AutoMigrate(models.Task{})
	if err != nil {
		log.Fatal("Failed to migrate database")
	}
	config.DB = db
	e := echo.New()
	e.GET("/tasks", handlers.GetTasks)
	e.POST("/tasks", handlers.PostTasks)
	e.PUT("/tasks/:id", handlers.UpdateTasks)
	err = e.Start(":8080")
	if err != nil {
		log.Fatal("Error starting server:")
	}
}
