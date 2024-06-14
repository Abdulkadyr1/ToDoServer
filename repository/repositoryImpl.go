package repository

import (
	"ToDo/config"
	"ToDo/tasks"
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

//curl -X POST -H "Content-Type: application/json"
//-d '{"title": "Sample Task", "description": "This is a sample task"}'
//localhost:8080/tasks
