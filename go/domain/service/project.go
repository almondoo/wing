package service

import (
	"wing/domain/entity"
	"wing/domain/repository"
	"wing/interface/validation"
)

type ProjectService interface {
	Get() ([]*entity.Project, error)
	GetDetail(id uint32) (*entity.Project, error)
	Create(*validation.ProjectRequest) error
	Update(uint32, *validation.ProjectRequest) error
	Delete(uint32) error
}

type projectService struct {
	projectRepo repository.ProjectRepository
}

func NewProjectService(projectRepo repository.ProjectRepository) ProjectService {
	return &projectService{projectRepo: projectRepo}
}

func (rs *projectService) Get() ([]*entity.Project, error) {
	return rs.projectRepo.Finds()
}
func (rs *projectService) GetDetail(id uint32) (*entity.Project, error) {
	return rs.projectRepo.FindByID(id)
}

func (rs *projectService) Create(request *validation.ProjectRequest) error {
	project := &entity.Project{
		Name:    request.Name,
		Content: request.Content,
	}
	_, err := rs.projectRepo.Create(project)
	return err
}

func (rs *projectService) Update(id uint32, request *validation.ProjectRequest) (err error) {
	var project *entity.Project
	project, err = rs.projectRepo.FindByID(id)
	if err != nil {
		return
	}
	project.Name = request.Name
	project.Content = request.Content
	_, err = rs.projectRepo.Update(project)
	return
}

func (rs *projectService) Delete(id uint32) (err error) {
	return rs.projectRepo.Delete(id)
}
