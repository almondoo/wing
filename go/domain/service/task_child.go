package service

import (
	"wing/domain/entity"
	"wing/domain/repository"
	"wing/interface/validation"
)

type TaskChildService interface {
	Get() ([]*entity.TaskChild, error)
	GetDetail(id uint) (*entity.TaskChild, error)
	Create(*validation.TaskChildRequest) error
	Update(uint, *validation.TaskChildRequest) error
	Delete(uint) error
}

type taskChildService struct {
	taskChildRepo repository.TaskChildRepository
}

func NewTaskChildService(taskChildRepo repository.TaskChildRepository) TaskChildService {
	return &taskChildService{taskChildRepo: taskChildRepo}
}

func (rs *taskChildService) Get() ([]*entity.TaskChild, error) {
	return rs.taskChildRepo.Finds()
}
func (rs *taskChildService) GetDetail(id uint) (*entity.TaskChild, error) {
	return rs.taskChildRepo.FindByID(id)
}

func (rs *taskChildService) Create(request *validation.TaskChildRequest) error {
	taskChild := &entity.TaskChild{}
	_, err := rs.taskChildRepo.Create(taskChild)
	return err
}

func (rs *taskChildService) Update(id uint, request *validation.TaskChildRequest) (err error) {
	var taskChild *entity.TaskChild
	taskChild, err = rs.taskChildRepo.FindByID(id)
	if err != nil {
		return
	}
	_, err = rs.taskChildRepo.Update(taskChild)
	return
}

func (rs *taskChildService) Delete(id uint) (err error) {
	return rs.taskChildRepo.Delete(id)
}
