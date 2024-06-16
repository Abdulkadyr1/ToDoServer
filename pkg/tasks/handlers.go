package tasks

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type TasksHandler struct {
	service ServiceTask
}

func (handler *TasksHandler) GetAllTasks(c echo.Context) error {
	result, err := handler.service.GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error getting tasks",
		})
	}
	return c.JSON(http.StatusOK, result)
}

func (handler *TasksHandler) PostTask(c echo.Context) error {
	task := &Task{}
	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error parsing task",
		})
	}
	err := handler.service.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error posting tasks",
		})
	}
	return c.JSON(http.StatusCreated, task)
}

func (handler *TasksHandler) UpdateTask(c echo.Context) error {
	task := &Task{}
	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error parsing task",
		})
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error parsing task id",
		})
	}
	err = handler.service.UpdateTask(task, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error updating tasks",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "OK",
	})
}

func (handler *TasksHandler) DeleteTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error parsing task id",
		})
	}
	err = handler.service.DeleteTask(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error deleting tasks",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DELETE SUCCESSFUL",
	})
}
