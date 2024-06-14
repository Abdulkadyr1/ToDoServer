package repository

import (
	"github.com/labstack/echo/v4"
)

type TaskRepository interface {
	GetTask(c echo.Context) error
}
