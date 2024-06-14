package handlers

import (
	"ToDo/config"
	"ToDo/repository"
	"ToDo/tasks"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type TasksHandler struct {
	repo repository.TaskRepositoryImpl
}

func (task *TasksHandler) GetAll(c echo.Context) error {
	err := task.repo.GetTasks(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error getting tasks",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "OK",
	})
}

func PostTasks(c echo.Context) error {
	tasks := new(tasks.Task)
	if c.Bind(tasks) != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": "Invalid request body",
		})
	}
	fmt.Printf("%s", tasks)
	//newTask := tasks.Task{Id: 11, Title: "aaa", Description: "bbb", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	result := config.DB.Create(tasks)
	if result.Error != nil {
		log.Fatal("Failed to create new task")
	}
	return nil
}

func UpdateTasks(c echo.Context) error {
	id := c.Param("id")
	var task tasks.Task
	result := config.DB.First(&task, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to find task",
		})
	}
	if c.Bind(&task) != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": "Invalid request body",
		})
	}
	config.DB.Save(&task)
	return nil
}

func DeleteTasks(c echo.Context) error {
	id := c.Param("id")
	var task tasks.Task
	result := config.DB.First(&task, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to find task",
		})
	}
	config.DB.Delete(&task)
	return nil
}
