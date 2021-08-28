package service

import (
	"wing/domain/entity"
	"wing/domain/repository"
	"wing/interface/validation"
)

type TaskPriorityService interface {
	Get() ([]*entity.TaskPriority, error)
	GetDetail(id uint) (*entity.TaskPriority, error)
	Create(*validation.TaskPriorityRequest) error
	Update(uint, *validation.TaskPriorityRequest) error
	Delete(uint) error
}

type taskPriorityService struct {
	taskPriorityRepo repository.TaskPriorityRepository
}

func NewTaskPriorityService(taskPriorityRepo repository.TaskPriorityRepository) TaskPriorityService {
	return &taskPriorityService{taskPriorityRepo: taskPriorityRepo}
}

func (rs *taskPriorityService) Get() ([]*entity.TaskPriority, error) {
	return rs.taskPriorityRepo.Finds()
}
func (rs *taskPriorityService) GetDetail(id uint) (*entity.TaskPriority, error) {
	return rs.taskPriorityRepo.FindByID(id)
}

func (rs *taskPriorityService) Create(request *validation.TaskPriorityRequest) error {
	taskPriority := &entity.TaskPriority{
		Name: request.Name,
	}
	_, err := rs.taskPriorityRepo.Create(taskPriority)
	return err
}

func (rs *taskPriorityService) Update(id uint, request *validation.TaskPriorityRequest) (err error) {
	var taskPriority *entity.TaskPriority
	taskPriority, err = rs.taskPriorityRepo.FindByID(id)
	if err != nil {
		return
	}
	_, err = rs.taskPriorityRepo.Update(taskPriority)
	return
}

func (rs *taskPriorityService) Delete(id uint) (err error) {
	return rs.taskPriorityRepo.Delete(id)
}
