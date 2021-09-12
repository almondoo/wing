package usecase

import (
	"wing/domain/service"
)

type UtilUsecase interface{}

type utilUsecase struct {
	us service.UtilService
}

func NewUtilUsecase(us service.UtilService) UtilUsecase {
	return &utilUsecase{us: us}
}
