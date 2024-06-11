package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Id          int
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
