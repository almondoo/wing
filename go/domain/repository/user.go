package repository

import "wing/domain/entity"

type UserRepository interface {
	FindByEmail(string) (*entity.User, error)
	FindByID(uint) (*entity.User, error)
	CreateUser(*entity.User) (*entity.User, error)
	Update(*entity.User) (*entity.User, error)
	Delete(*entity.User) error
}
