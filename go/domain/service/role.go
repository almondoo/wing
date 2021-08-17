package service

import (
	"wing/domain/entity"
	"wing/domain/repository"
)

type RoleService interface {
	Create(product *entity.Role) (string, error)
}

type roleService struct {
	roleRepo repository.RoleRepository
}

func NewRoleService(roleRepo repository.RoleRepository) RoleService {
	return &roleService{roleRepo: roleRepo}
}

func (as *roleService) Create(product *entity.Role) (string, error) {
	return "", nil
}
