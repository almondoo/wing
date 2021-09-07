package repository

import "wing/domain/entity"

type RoleRepository interface {
	FindByID(roleId uint32) (*entity.Role, error)
	Finds() ([]*entity.Role, error)
	FindByConditions(map[string]interface{}) (*entity.Role, error)
	Create(*entity.Role) (*entity.Role, error)
	Update(*entity.Role) (*entity.Role, error)
	Delete(roleId uint32) error
}
