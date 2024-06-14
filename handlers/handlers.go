package handlers

import (
	"ToDo/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TasksHandler struct {
	repo repository.TaskRepositoryImpl
}

func (handler *TasksHandler) GetAll(c echo.Context) error {
	err := handler.repo.GetTasks(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error getting tasks",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "OK",
	})
}

func (handler *TasksHandler) Post(c echo.Context) error {
	err := handler.repo.PostTasks(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error posting tasks",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "OK",
	})
}

func (handler *TasksHandler) Update(c echo.Context) error {
	err := handler.repo.UpdateTasks(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error updating tasks",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "OK",
	})
}

func (handler *TasksHandler) Delete(c echo.Context) error {
	err := handler.repo.DeleteTasks(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error deleting tasks",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "OK",
	})
}
