package tasks

import (
	"ToDo/config"
)

type TaskRepository interface {
	GetAll() ([]Task, error)
	Create(task *Task) error
	Update(task *Task, id int) error
	Delete(id int) error
}

type TaskRepositoryImpl struct{}

func (a *TaskRepositoryImpl) GetAll() ([]Task, error) {
	var tasksArr []Task
	err := config.DB.Find(&tasksArr).Error
	if err != nil {
		return nil, err
	}
	return tasksArr, nil
}

func (a *TaskRepositoryImpl) Create(task *Task) error {
	result := config.DB.Create(task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *TaskRepositoryImpl) Update(task *Task, id int) error {
	var existingTask Task
	result := config.DB.First(&existingTask, id)
	if result.Error != nil {
		return result.Error
	}

	// Обновление полей на основе значений из task
	if err := config.DB.Model(&existingTask).Updates(task).Error; err != nil {
		return err
	}

	return nil
}

func (a *TaskRepositoryImpl) Delete(id int) error {
	return config.DB.Delete(Task{}, id).Error
}

//curl -X POST -H "Content-Type: application/json" -d '{"title": "Sample Task", "description": "This is a sample task"}' localhost:8080/tasks
