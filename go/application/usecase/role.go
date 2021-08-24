package usecase

import (
	"wing/domain/service"
	"wing/interface/validation"
)

type RoleUsecase interface {
	Create(*validation.RoleCreateRequest) error
	Update(*validation.RoleUpdateRequest) error
}

type roleUsecase struct {
	rs service.RoleService
}

func NewRoleUsecase(rs service.RoleService) RoleUsecase {
	return &roleUsecase{rs: rs}
}

func (ru *roleUsecase) Create(request *validation.RoleCreateRequest) error {
	return ru.rs.Create(request)
}

func (ru *roleUsecase) Update(request *validation.RoleUpdateRequest) error {
	return ru.rs.Update(request)
}
