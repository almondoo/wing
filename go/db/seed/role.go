package seed

import (
	"wing/domain/entity"

	"gorm.io/gorm"
)

type roleSeeder struct {
	conn *gorm.DB
}

func NewRoleSeeder(db *gorm.DB) *roleSeeder {
	return &roleSeeder{conn: db}
}

func (rs *roleSeeder) Seeder() {
	rs.create(entity.Role{
		ID:   1,
		Name: "開発者",
	})
	rs.create(entity.Role{
		ID:   5,
		Name: "統括者",
	})
	rs.create(entity.Role{
		ID:   10,
		Name: "管理者",
	})
	rs.create(entity.Role{
		ID:   20,
		Name: "閲覧者",
	})
}

func (rs *roleSeeder) create(entity entity.Role) error {
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
