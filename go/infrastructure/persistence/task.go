package persistence

import (
	"wing/domain/entity"
	"wing/domain/repository"

	"gorm.io/gorm"
)

type taskRepository struct {
	Conn *gorm.DB
}

func NewTaskRepository(conn *gorm.DB) repository.TaskRepository {
	return &taskRepository{Conn: conn}
}

// IDで取得
func (rr *taskRepository) FindByID(id uint) (*entity.Task, error) {
	task := &entity.Task{ID: id}

	if err := rr.Conn.First(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (rr *taskRepository) Finds() ([]*entity.Task, error) {
	var tasks []*entity.Task
	if err := rr.Conn.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// Create 作成
func (rr *taskRepository) Create(task *entity.Task) (*entity.Task, error) {
	tx := rr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Create(&task).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return task, tx.Commit().Error
}

// Update 更新
func (rr *taskRepository) Update(task *entity.Task) (*entity.Task, error) {
	tx := rr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Model(&task).Updates(&task).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return task, tx.Commit().Error
}

// Delete 削除
func (rr *taskRepository) Delete(id uint) error {
	tx := rr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Delete(&entity.Task{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
