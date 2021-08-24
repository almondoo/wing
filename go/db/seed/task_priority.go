package seed

import (
	"wing/domain/entity"

	"gorm.io/gorm"
)

type taskPrioritySeeder struct {
	conn *gorm.DB
}

func NewTaskPrioritySeeder(db *gorm.DB) *taskPrioritySeeder {
	return &taskPrioritySeeder{conn: db}
}

func (tps *taskPrioritySeeder) Seeder() {
	tps.create(entity.TaskPriority{
		ID:   1,
		Name: "低優先",
	})
	tps.create(entity.TaskPriority{
		ID:   2,
		Name: "中優先",
	})
	tps.create(entity.TaskPriority{
		ID:   3,
		Name: "高優先",
	})
}

func (tps *taskPrioritySeeder) create(entity entity.TaskPriority) error {
	tx := tps.conn.Begin()
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
