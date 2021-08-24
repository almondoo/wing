package usecase

import "wing/domain/service"

type SampleUsecase interface {
	Logic() error
}

type sampleUsecase struct {
	s service.SampleService
}

func NewSampleUsecase(s service.SampleService) SampleUsecase {
	return &sampleUsecase{s: s}
}

func (su *sampleUsecase) Logic() error {
	return nil
}
