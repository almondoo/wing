package usecase

import (
	"wing/domain/entity"
	"wing/domain/service"
	"wing/interface/validation"
)

type RoleUsecase interface {
	Get() ([]*entity.Role, error)
	GetDetail(id uint) (*entity.Role, error)
	Create(*validation.RoleRequest) error
	Update(uint, *validation.RoleRequest) error
	Delete(uint) error
}

type roleUsecase struct {
	rs service.RoleService
}

func NewRoleUsecase(rs service.RoleService) RoleUsecase {
	return &roleUsecase{rs: rs}
}

func (ru *roleUsecase) Get() ([]*entity.Role, error) {
	return ru.rs.Get()
}
func (ru *roleUsecase) GetDetail(id uint) (*entity.Role, error) {
	return ru.rs.GetDetail(id)
}

func (ru *roleUsecase) Create(request *validation.RoleRequest) error {
	return ru.rs.Create(request)
}

func (ru *roleUsecase) Update(id uint, request *validation.RoleRequest) error {
	return ru.rs.Update(id, request)
}

func (ru *roleUsecase) Delete(id uint) error {
	return ru.rs.Delete(id)
}
