package usecase

import (
	"wing/domain/entity"
	"wing/domain/service"
	"wing/interface/validation"
)

type TaskUsecase interface {
	Get() ([]*entity.Task, error)
	GetDetail(id uint) (*entity.Task, error)
	Create(*validation.TaskRequest) error
	Update(uint, *validation.TaskRequest) error
	Delete(uint) error
}

type taskUsecase struct {
	ts service.TaskService
}

func NewTaskUsecase(ts service.TaskService) TaskUsecase {
	return &taskUsecase{ts: ts}
}

func (tu *taskUsecase) Get() ([]*entity.Task, error) {
	return tu.ts.Get()
}
func (tu *taskUsecase) GetDetail(id uint) (*entity.Task, error) {
	return tu.ts.GetDetail(id)
}

func (tu *taskUsecase) Create(request *validation.TaskRequest) error {
	return tu.ts.Create(request)
}

func (tu *taskUsecase) Update(id uint, request *validation.TaskRequest) error {
	return tu.ts.Update(id, request)
}

func (tu *taskUsecase) Delete(id uint) error {
	return tu.ts.Delete(id)
}
