package service

import "wing/domain/repository"

type SampleService interface {
	Sample() error
}

type sampleService struct {
	sampleRepo repository.SampleRepository
}

func NewSampleService(sampleRepo repository.SampleRepository) SampleService {
	return &sampleService{sampleRepo: sampleRepo}
}

func (us *sampleService) Sample() error {
	// ここにビジネスロジックを書く
	return nil
}
