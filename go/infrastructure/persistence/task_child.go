package persistence

import (
	"wing/domain/entity"
	"wing/domain/repository"

	"gorm.io/gorm"
)

type taskChildRepository struct {
	Conn *gorm.DB
}

func NewTaskChildRepository(conn *gorm.DB) repository.TaskChildRepository {
	return &taskChildRepository{Conn: conn}
}

// IDで取得
func (rr *taskChildRepository) FindByID(id uint) (*entity.TaskChild, error) {
	taskChild := &entity.TaskChild{ID: id}

	if err := rr.Conn.First(&taskChild).Error; err != nil {
		return nil, err
	}

	return taskChild, nil
}

func (rr *taskChildRepository) Finds() ([]*entity.TaskChild, error) {
	var taskChilds []*entity.TaskChild
	if err := rr.Conn.Find(&taskChilds).Error; err != nil {
		return nil, err
	}
	return taskChilds, nil
}

// Create 作成
func (rr *taskChildRepository) Create(taskChild *entity.TaskChild) (*entity.TaskChild, error) {
	tx := rr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Create(&taskChild).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return taskChild, tx.Commit().Error
}

// Update 更新
func (rr *taskChildRepository) Update(taskChild *entity.TaskChild) (*entity.TaskChild, error) {
	tx := rr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Model(&taskChild).Updates(&taskChild).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return taskChild, tx.Commit().Error
}

// Delete 削除
func (rr *taskChildRepository) Delete(id uint) error {
	tx := rr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Delete(&entity.TaskChild{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
