package usecase

import (
	"wing/domain/entity"
	"wing/domain/service"
	"wing/interface/validation"
)

type TaskStatusUsecase interface {
	Get() ([]*entity.TaskStatus, error)
	GetDetail(id uint) (*entity.TaskStatus, error)
	Create(*validation.TaskStatusRequest) error
	Update(uint, *validation.TaskStatusRequest) error
	Delete(uint) error
}

type taskStatusUsecase struct {
	tss service.TaskStatusService
}

func NewTaskStatusUsecase(tss service.TaskStatusService) TaskStatusUsecase {
	return &taskStatusUsecase{tss: tss}
}

func (tsu *taskStatusUsecase) Get() ([]*entity.TaskStatus, error) {
	return tsu.tss.Get()
}
func (tsu *taskStatusUsecase) GetDetail(id uint) (*entity.TaskStatus, error) {
	return tsu.tss.GetDetail(id)
}

func (tsu *taskStatusUsecase) Create(request *validation.TaskStatusRequest) error {
	return tsu.tss.Create(request)
}

func (tsu *taskStatusUsecase) Update(id uint, request *validation.TaskStatusRequest) error {
	return tsu.tss.Update(id, request)
}

func (tsu *taskStatusUsecase) Delete(id uint) error {
	return tsu.tss.Delete(id)
}
