package repository

import "wing/domain/entity"

type ProjectRepository interface {
	FindByID(uint32) (*entity.Project, error)
	Finds() ([]*entity.Project, error)
	Create(*entity.Project) (*entity.Project, error)
	Update(*entity.Project) (*entity.Project, error)
	Delete(uint32) error
}
