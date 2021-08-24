package seed

import (
	"wing/domain/entity"

	"gorm.io/gorm"
)

type taskStatusSeeder struct {
	conn *gorm.DB
}

func NewTaskStatusSeeder(db *gorm.DB) *taskStatusSeeder {
	return &taskStatusSeeder{conn: db}
}

func (tss *taskStatusSeeder) Seeder() {
	tss.create(entity.TaskStatus{
		ID:   1,
		Name: "新規",
	})
	tss.create(entity.TaskStatus{
		ID:   2,
		Name: "進行中",
	})
	tss.create(entity.TaskStatus{
		ID:   3,
		Name: "レビュー待ち",
	})
	tss.create(entity.TaskStatus{
		ID:   4,
		Name: "完了",
	})
	tss.create(entity.TaskStatus{
		ID:   5,
		Name: "キャンセル",
	})
}

func (tss *taskStatusSeeder) create(entity entity.TaskStatus) error {
	tx := tss.conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&entity).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
