package main

import (
	"ToDo/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {
	//e := echo.New()
	//e.GET("/tasks", handlers.GetTasks)
	//e.Start(":8080")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	err = db.AutoMigrate(models.Task{})
	if err != nil {
		log.Fatal("Failed to migrate database")
	}
	newTask := models.Task{Id: 11, Title: "aaa", Description: "bbb", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	result := db.Create(&newTask)
	if result.Error != nil {
		log.Fatal("Faile to create new task")
	}
	var tasks []models.Task
	db.Find(&tasks)
	fmt.Println(tasks)
}
