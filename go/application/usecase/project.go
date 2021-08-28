package usecase

import (
	"wing/domain/entity"
	"wing/domain/service"
	"wing/interface/validation"
)

type ProjectUsecase interface {
	Get() ([]*entity.Project, error)
	GetDetail(id uint32) (*entity.Project, error)
	Create(*validation.ProjectRequest) error
	Update(uint32, *validation.ProjectRequest) error
	Delete(uint32) error
}

type projectUsecase struct {
	ps service.ProjectService
}

func NewProjectUsecase(ps service.ProjectService) ProjectUsecase {
	return &projectUsecase{ps: ps}
}

func (tsu *projectUsecase) Get() ([]*entity.Project, error) {
	return tsu.ps.Get()
}
func (tsu *projectUsecase) GetDetail(id uint32) (*entity.Project, error) {
	return tsu.ps.GetDetail(id)
}

func (tsu *projectUsecase) Create(request *validation.ProjectRequest) error {
	return tsu.ps.Create(request)
}

func (tsu *projectUsecase) Update(id uint32, request *validation.ProjectRequest) error {
	return tsu.ps.Update(id, request)
}

func (tsu *projectUsecase) Delete(id uint32) error {
	return tsu.ps.Delete(id)
}
