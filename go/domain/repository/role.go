package repository

import "wing/domain/entity"

type RoleRepository interface {
	CreateRole(role *entity.Role) error
	FindByEmail(email string) (entity.Role, error)
	Create(role *entity.Role) (*entity.Role, error)
	FindByID(id uint) (*entity.Role, error)
	Update(role *entity.Role) (*entity.Role, error)
	Delete(role *entity.Role) error
}
