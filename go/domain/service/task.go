package service

import (
	"wing/domain/entity"
	"wing/domain/repository"
	"wing/interface/validation"
)

type TaskService interface {
	Get() ([]*entity.Task, error)
	GetDetail(id uint) (*entity.Task, error)
	Create(*validation.TaskRequest) error
	Update(uint, *validation.TaskRequest) error
	Delete(uint) error
}

type taskService struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(taskRepo repository.TaskRepository) TaskService {
	return &taskService{taskRepo: taskRepo}
}

func (rs *taskService) Get() ([]*entity.Task, error) {
	return rs.taskRepo.Finds()
}
func (rs *taskService) GetDetail(id uint) (*entity.Task, error) {
	return rs.taskRepo.FindByID(id)
}

func (rs *taskService) Create(request *validation.TaskRequest) error {
	task := &entity.Task{}
	_, err := rs.taskRepo.Create(task)
	return err
}

func (rs *taskService) Update(id uint, request *validation.TaskRequest) (err error) {
	var task *entity.Task
	task, err = rs.taskRepo.FindByID(id)
	if err != nil {
		return
	}
	_, err = rs.taskRepo.Update(task)
	return
}

func (rs *taskService) Delete(id uint) (err error) {
	return rs.taskRepo.Delete(id)
}
