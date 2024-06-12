package handlers

import (
	"ToDo/config"
	"ToDo/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func GetTasks(c echo.Context) error {
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	var tasks []models.Task
	db.Find(&tasks)
	fmt.Println(tasks)
	return nil
}

func PostTasks(c echo.Context) error {
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
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
		log.Fatal("Failed to create new task")
	}
	var tasks []models.Task
	db.Find(&tasks)
	fmt.Println(tasks)
	return nil
}
