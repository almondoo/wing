package persistence

import (
	"wing/domain/entity"
	"wing/domain/repository"

	"gorm.io/gorm"
)

type taskStatusRepository struct {
	Conn *gorm.DB
}

func NewTaskStatusRepository(conn *gorm.DB) repository.TaskStatusRepository {
	return &taskStatusRepository{Conn: conn}
}

// IDで取得
func (tsr *taskStatusRepository) FindByID(id uint) (*entity.TaskStatus, error) {
	role := &entity.TaskStatus{ID: id}

	if err := tsr.Conn.First(&role).Error; err != nil {
		return nil, err
	}

	return role, nil
}

func (tsr *taskStatusRepository) Finds() ([]*entity.TaskStatus, error) {
	var roles []*entity.TaskStatus
	if err := tsr.Conn.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

// 作成
func (tsr *taskStatusRepository) Create(taskStatus *entity.TaskStatus) (*entity.TaskStatus, error) {
	tx := tsr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Create(&taskStatus).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return taskStatus, tx.Commit().Error
}

// 更新
func (tsr *taskStatusRepository) Update(taskStatus *entity.TaskStatus) (*entity.TaskStatus, error) {
	tx := tsr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Model(&taskStatus).Updates(&taskStatus).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return taskStatus, tx.Commit().Error
}

// 削除
func (tsr *taskStatusRepository) Delete(id uint) error {
	tx := tsr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Delete(&entity.TaskStatus{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
