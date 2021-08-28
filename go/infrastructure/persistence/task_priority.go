package persistence

import (
	"wing/domain/entity"
	"wing/domain/repository"

	"gorm.io/gorm"
)

type taskPriorityRepository struct {
	Conn *gorm.DB
}

func NewTaskPriorityRepository(conn *gorm.DB) repository.TaskPriorityRepository {
	return &taskPriorityRepository{Conn: conn}
}

// IDで取得
func (tsr *taskPriorityRepository) FindByID(id uint) (*entity.TaskPriority, error) {
	role := &entity.TaskPriority{ID: id}

	if err := tsr.Conn.First(&role).Error; err != nil {
		return nil, err
	}

	return role, nil
}

func (tsr *taskPriorityRepository) Finds() ([]*entity.TaskPriority, error) {
	var roles []*entity.TaskPriority
	if err := tsr.Conn.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

// 作成
func (tsr *taskPriorityRepository) Create(taskPriority *entity.TaskPriority) (*entity.TaskPriority, error) {
	tx := tsr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Create(&taskPriority).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return taskPriority, tx.Commit().Error
}

// 更新
func (tsr *taskPriorityRepository) Update(taskPriority *entity.TaskPriority) (*entity.TaskPriority, error) {
	tx := tsr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Model(&taskPriority).Updates(&taskPriority).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return taskPriority, tx.Commit().Error
}

// 削除
func (tsr *taskPriorityRepository) Delete(id uint) error {
	tx := tsr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Delete(&entity.TaskPriority{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
