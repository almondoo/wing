package persistence

import (
	"wing/domain/repository"

	"gorm.io/gorm"
)

type TaskRepository struct {
	Conn *gorm.DB
}

func NewTaskRepository(conn *gorm.DB) repository.TaskRepository {
	return &TaskRepository{Conn: conn}
}
