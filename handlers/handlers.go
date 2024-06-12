package handlers

import (
	"ToDo/config"
	"ToDo/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func GetTasks(c echo.Context) error {
	var tasks []models.Task
	result := config.DB.Find(&tasks)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve tasks",
		})
	}
	return c.JSON(http.StatusOK, tasks)
}

func PostTasks(c echo.Context) error {
	tasks := new(models.Task)
	if c.Bind(tasks) != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": "Invalid request body",
		})
	}
	fmt.Print(c)
	//newTask := models.Task{Id: 11, Title: "aaa", Description: "bbb", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	result := config.DB.Create(tasks)
	if result.Error != nil {
		log.Fatal("Failed to create new task")
	}
	return nil
}
