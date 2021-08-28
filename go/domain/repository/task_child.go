package repository

import "wing/domain/entity"

type TaskChildRepository interface {
	FindByID(id uint) (*entity.TaskChild, error)
	Finds() ([]*entity.TaskChild, error)
	Create(*entity.TaskChild) (*entity.TaskChild, error)
	Update(*entity.TaskChild) (*entity.TaskChild, error)
	Delete(uint) error
}
