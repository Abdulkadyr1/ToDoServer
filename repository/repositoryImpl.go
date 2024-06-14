package repository

import (
	"ToDo/config"
	"ToDo/tasks"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TaskRepositoryImpl struct{}

func (a *TaskRepositoryImpl) GetTasks(c echo.Context) error {
	var tasksArr []tasks.Task
	result := config.DB.Find(&tasksArr)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to retrieve tasksArr",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"tasks": tasksArr,
	})
}

func (a *TaskRepositoryImpl) PostTasks(c echo.Context) error {
	tasks := new(tasks.Task)
	if c.Bind(tasks) != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": "Invalid request body",
		})
	}
	fmt.Printf("%s", tasks)
	result := config.DB.Create(tasks)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to create tasks",
		})
	}
	return nil
}

func (a *TaskRepositoryImpl) UpdateTasks(c echo.Context) error {
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

func (a *TaskRepositoryImpl) DeleteTasks(c echo.Context) error {
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

//curl -X POST -H "Content-Type: application/json"
//-d '{"title": "Sample Task", "description": "This is a sample task"}'
//localhost:8080/tasks
