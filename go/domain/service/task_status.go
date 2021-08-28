package service

import (
	"wing/domain/entity"
	"wing/domain/repository"
	"wing/interface/validation"
)

type TaskStatusService interface {
	Get() ([]*entity.TaskStatus, error)
	GetDetail(id uint) (*entity.TaskStatus, error)
	Create(*validation.TaskStatusRequest) error
	Update(uint, *validation.TaskStatusRequest) error
	Delete(uint) error
}

type taskStatusService struct {
	taskStatusRepo repository.TaskStatusRepository
}

func NewTaskStatusService(taskStatusRepo repository.TaskStatusRepository) TaskStatusService {
	return &taskStatusService{taskStatusRepo: taskStatusRepo}
}

func (rs *taskStatusService) Get() ([]*entity.TaskStatus, error) {
	return rs.taskStatusRepo.Finds()
}
func (rs *taskStatusService) GetDetail(id uint) (*entity.TaskStatus, error) {
	return rs.taskStatusRepo.FindByID(id)
}

func (rs *taskStatusService) Create(request *validation.TaskStatusRequest) error {
	taskStatus := &entity.TaskStatus{
		Name: request.Name,
	}
	_, err := rs.taskStatusRepo.Create(taskStatus)
	return err
}

func (rs *taskStatusService) Update(id uint, request *validation.TaskStatusRequest) (err error) {
	var taskStatus *entity.TaskStatus
	taskStatus, err = rs.taskStatusRepo.FindByID(id)
	if err != nil {
		return
	}
	_, err = rs.taskStatusRepo.Update(taskStatus)
	return
}

func (rs *taskStatusService) Delete(id uint) (err error) {
	return rs.taskStatusRepo.Delete(id)
}
