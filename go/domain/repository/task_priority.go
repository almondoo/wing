package repository

import "wing/domain/entity"

type TaskPriorityRepository interface {
	Finds() ([]*entity.TaskPriority, error)
	FindByID(uint) (*entity.TaskPriority, error)
	Create(*entity.TaskPriority) (*entity.TaskPriority, error)
	Update(*entity.TaskPriority) (*entity.TaskPriority, error)
	Delete(uint) error
}
