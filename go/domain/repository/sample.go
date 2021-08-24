package repository

type SampleRepository interface {
	Create() error
	Update() error
	Delete() error
}
