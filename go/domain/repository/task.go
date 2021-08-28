package repository

import "wing/domain/entity"

type TaskRepository interface {
	FindByID(id uint) (*entity.Task, error)
	Finds() ([]*entity.Task, error)
	Create(*entity.Task) (*entity.Task, error)
	Update(*entity.Task) (*entity.Task, error)
	Delete(uint) error
}
