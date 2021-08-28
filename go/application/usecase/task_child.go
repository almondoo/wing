package usecase

import (
	"wing/domain/entity"
	"wing/domain/service"
	"wing/interface/validation"
)

type TaskChildUsecase interface {
	Get() ([]*entity.TaskChild, error)
	GetDetail(id uint) (*entity.TaskChild, error)
	Create(*validation.TaskChildRequest) error
	Update(uint, *validation.TaskChildRequest) error
	Delete(uint) error
}

type taskChildUsecase struct {
	tcs service.TaskChildService
}

func NewTaskChildUsecase(tcs service.TaskChildService) TaskChildUsecase {
	return &taskChildUsecase{tcs: tcs}
}

func (tcu *taskChildUsecase) Get() ([]*entity.TaskChild, error) {
	return tcu.tcs.Get()
}
func (tcu *taskChildUsecase) GetDetail(id uint) (*entity.TaskChild, error) {
	return tcu.tcs.GetDetail(id)
}

func (tcu *taskChildUsecase) Create(request *validation.TaskChildRequest) error {
	return tcu.tcs.Create(request)
}

func (tcu *taskChildUsecase) Update(id uint, request *validation.TaskChildRequest) error {
	return tcu.tcs.Update(id, request)
}

func (tcu *taskChildUsecase) Delete(id uint) error {
	return tcu.tcs.Delete(id)
}
