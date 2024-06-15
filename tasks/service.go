package tasks

type Service interface {
	GetAllTasks() ([]Task, error)
	CreateTask(task *Task) error
	UpdateTask(task *Task, id int) error
	DeleteTask(id int) error
}

type ServiceTask struct {
	repo TaskRepositoryImpl
}

func (service *ServiceTask) GetAllTasks() ([]Task, error) {
	return service.repo.GetAll()
}

func (service *ServiceTask) CreateTask(task *Task) error {
	return service.repo.Create(task)
}

func (service *ServiceTask) UpdateTask(task *Task, id int) error {
	return service.repo.Update(task, id)
}

func (service *ServiceTask) DeleteTask(id int) error {
	return service.repo.Delete(id)
}
