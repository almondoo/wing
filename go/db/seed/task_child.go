package seed

import (
	"wing/domain/entity"

	"gorm.io/gorm"
)

type taskChildSeeder struct {
	conn *gorm.DB
}

func NewTaskChildSeeder(db *gorm.DB) *taskChildSeeder {
	return &taskChildSeeder{conn: db}
}

func (ts *taskChildSeeder) Seeder() {
	ts.create(entity.TaskChild{
		ID:             1,
		TaskID:         1,
		CreateUserID:   1,
		Title:          "テスト",
		Content:        "テストで作ったよ。",
		TaskStatusID:   1,
		AssignUserID:   1,
		TaskPriorityID: 1,
	})
}

func (ts *taskChildSeeder) create(entity entity.TaskChild) error {
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
