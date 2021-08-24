package service

import (
	"wing/domain/entity"
	"wing/domain/repository"
	"wing/interface/validation"
)

type RoleService interface {
	Create(*validation.RoleCreateRequest) error
	Update(*validation.RoleUpdateRequest) error
	Delete(*validation.RoleDeleteRequest) error
}

type roleService struct {
	roleRepo repository.RoleRepository
}

func NewRoleService(roleRepo repository.RoleRepository) RoleService {
	return &roleService{roleRepo: roleRepo}
}

func (rs *roleService) Create(request *validation.RoleCreateRequest) error {
	role := &entity.Role{
		Name: request.Name,
	}
	_, err := rs.roleRepo.Create(role)
	return err
}

func (rs *roleService) Update(request *validation.RoleUpdateRequest) (err error) {
	var role *entity.Role
	role, err = rs.roleRepo.FindByID(request.ID)
	if err != nil {
		return
	}
	_, err = rs.roleRepo.Update(role)
	return
}

func (rs *roleService) Delete(request *validation.RoleDeleteRequest) (err error) {
	return rs.roleRepo.Delete(request.ID)
}
