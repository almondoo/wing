package repository

import "wing/domain/entity"

type RoleRepository interface {
	FindByID(id uint) (*entity.Role, error)
	Finds() ([]*entity.Role, error)
	Create(*entity.Role) (*entity.Role, error)
	Update(*entity.Role) (*entity.Role, error)
	Delete(uint) error
}
