package repository

import "wing/domain/entity"

type TaskStatusRepository interface {
	Finds() ([]*entity.TaskStatus, error)
	FindByID(uint) (*entity.TaskStatus, error)
	Create(*entity.TaskStatus) (*entity.TaskStatus, error)
	Update(*entity.TaskStatus) (*entity.TaskStatus, error)
	Delete(uint) error
}
