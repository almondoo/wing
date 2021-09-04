package service

import (
	"wing/domain/entity"
	"wing/domain/repository"
)

type UtilService interface {
	FindUser(uint) (*entity.User, error)
}

type utilService struct {
	userRepo repository.UserRepository
	roleRepo repository.RoleRepository
}

func NewUtilService(userRepo repository.UserRepository, roleRepo repository.RoleRepository) UtilService {
	return &utilService{userRepo: userRepo, roleRepo: roleRepo}
}

func (rs *utilService) FindUser(id uint) (*entity.User, error) {
	return rs.userRepo.FindByID(id)
}
