package usecase

import (
	"wing/domain/entity"
	"wing/domain/service"
	"wing/interface/validation"
)

type TaskPriorityUsecase interface {
	Get() ([]*entity.TaskPriority, error)
	GetDetail(id uint) (*entity.TaskPriority, error)
	Create(*validation.TaskPriorityRequest) error
	Update(uint, *validation.TaskPriorityRequest) error
	Delete(uint) error
}

type taskPriorityUsecase struct {
	tss service.TaskPriorityService
}

func NewTaskPriorityUsecase(tss service.TaskPriorityService) TaskPriorityUsecase {
	return &taskPriorityUsecase{tss: tss}
}

func (tsu *taskPriorityUsecase) Get() ([]*entity.TaskPriority, error) {
	return tsu.tss.Get()
}
func (tsu *taskPriorityUsecase) GetDetail(id uint) (*entity.TaskPriority, error) {
	return tsu.tss.GetDetail(id)
}

func (tsu *taskPriorityUsecase) Create(request *validation.TaskPriorityRequest) error {
	return tsu.tss.Create(request)
}

func (tsu *taskPriorityUsecase) Update(id uint, request *validation.TaskPriorityRequest) error {
	return tsu.tss.Update(id, request)
}

func (tsu *taskPriorityUsecase) Delete(id uint) error {
	return tsu.tss.Delete(id)
}
