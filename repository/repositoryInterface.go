package repository

import (
	"github.com/labstack/echo/v4"
)

type TaskRepository interface {
	GetTask(c echo.Context) error
	PostTasks(c echo.Context) error
	UpdateTasks(c echo.Context) error
	DeleteTasks(c echo.Context) error
}
