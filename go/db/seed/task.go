package seed

import (
	"wing/domain/entity"

	"gorm.io/gorm"
)

type taskSeeder struct {
	conn *gorm.DB
}

func NewTaskSeeder(db *gorm.DB) *taskSeeder {
	return &taskSeeder{conn: db}
}

func (ts *taskSeeder) Seeder() {
	ts.create(entity.Task{
		ID:             1,
		ProjectID:      1,
		CreateUserID:   1,
		Title:          "テスト",
		Content:        "テストで作ったよ。",
		TaskStatusID:   1,
		AssignUserID:   1,
		TaskPriorityID: 1,
	})
}

func (ts *taskSeeder) create(entity entity.Task) error {
	tx := ts.conn.Begin()
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
