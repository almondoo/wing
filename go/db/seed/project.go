package seed

import (
	"wing/domain/entity"

	"gorm.io/gorm"
)

type projectSeeder struct {
	conn *gorm.DB
}

func NewProjectSeeder(db *gorm.DB) *projectSeeder {
	return &projectSeeder{conn: db}
}

func (rs *projectSeeder) Seeder() {
	rs.create(entity.Project{
		ID:      1,
		Name:    "Wing",
		Content: "プロジェクトだよ。いいね。",
	})
}

func (rs *projectSeeder) create(entity entity.Project) error {
	tx := rs.conn.Begin()
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
