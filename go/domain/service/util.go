package service

import (
	"wing/domain/repository"
)

type UtilService interface{}

type utilService struct {
	userRepo repository.UserRepository
	roleRepo repository.RoleRepository
}

func NewUtilService(userRepo repository.UserRepository, roleRepo repository.RoleRepository) UtilService {
	return &utilService{userRepo: userRepo, roleRepo: roleRepo}
}
