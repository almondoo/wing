package service

import (
	"wing/domain/entity"
	"wing/domain/repository"
	"wing/interface/validation"
)

type RoleService interface {
	Get() ([]*entity.Role, error)
	GetDetail(id uint) (*entity.Role, error)
	Create(*validation.RoleRequest) error
	Update(uint, *validation.RoleRequest) error
	Delete(uint) error
}

type roleService struct {
	roleRepo repository.RoleRepository
}

func NewRoleService(roleRepo repository.RoleRepository) RoleService {
	return &roleService{roleRepo: roleRepo}
}

func (rs *roleService) Get() ([]*entity.Role, error) {
	return rs.roleRepo.Finds()
}
func (rs *roleService) GetDetail(id uint) (*entity.Role, error) {
	return rs.roleRepo.FindByID(id)
}

func (rs *roleService) Create(request *validation.RoleRequest) error {
	role := &entity.Role{
		Name: request.Name,
	}
	_, err := rs.roleRepo.Create(role)
	return err
}

func (rs *roleService) Update(id uint, request *validation.RoleRequest) (err error) {
	var role *entity.Role
	role, err = rs.roleRepo.FindByID(id)
	if err != nil {
		return
	}
	_, err = rs.roleRepo.Update(role)
	return
}

func (rs *roleService) Delete(id uint) (err error) {
	return rs.roleRepo.Delete(id)
}
