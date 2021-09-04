package usecase

import (
	"wing/domain/entity"
	"wing/domain/service"
)

type UtilUsecase interface {
	FindUser(uint) (*entity.User, error)
}

type utilUsecase struct {
	us service.UtilService
}

func NewUtilUsecase(us service.UtilService) UtilUsecase {
	return &utilUsecase{us: us}
}

func (uu *utilUsecase) FindUser(id uint) (*entity.User, error) {
	return uu.us.FindUser(id)
}
