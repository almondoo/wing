package repository

import "wing/domain/entity"

type UserHaveRoleRepository interface {
	FindsByUserID(userId uint) ([]*entity.UserHaveRole, error)
	FindByConditions(conditions map[string]interface{}) (*entity.UserHaveRole, error)
	Create(userHaveRole *entity.UserHaveRole) (*entity.UserHaveRole, error)
	Update(userHaveRole *entity.UserHaveRole) (*entity.UserHaveRole, error)
	Delete(userId uint, roleId uint) error
}
