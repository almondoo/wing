package usecase

import (
	"wing/domain/entity"
	"wing/domain/service"
	"wing/interface/validation"
)

type RoleUsecase interface {
	HasRole(userId uint, operation string) bool
	Get() ([]*entity.Role, error)
	GetDetail(id uint32) (*entity.Role, error)
	Create(*validation.RoleRequest) error
	Update(uint32, *validation.RoleRequest) error
	Delete(uint32) error
}

type roleUsecase struct {
	rs service.RoleService
	us service.UserService
}

func NewRoleUsecase(rs service.RoleService, us service.UserService) RoleUsecase {
	return &roleUsecase{rs: rs, us: us}
}

func (ru *roleUsecase) HasRole(userId uint, operation string) bool {
	user, err := ru.us.Find(userId)
	if err != nil {
		return false
	}
	return ru.rs.HasRole(user.RoleID, operation)
}

func (ru *roleUsecase) Get() ([]*entity.Role, error) {
	return ru.rs.Get()
}
func (ru *roleUsecase) GetDetail(id uint32) (*entity.Role, error) {
	return ru.rs.GetDetail(id)
}

func (ru *roleUsecase) Create(request *validation.RoleRequest) error {
	return ru.rs.Create(request)
}

func (ru *roleUsecase) Update(id uint32, request *validation.RoleRequest) error {
	return ru.rs.Update(id, request)
}

func (ru *roleUsecase) Delete(id uint32) error {
	return ru.rs.Delete(id)
}
