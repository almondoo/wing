package seed

import (
	"wing/domain/entity"
	"wing/infrastructure/security"

	"gorm.io/gorm"
)

type userSeeder struct {
	conn *gorm.DB
}

func NewUserSeeder(db *gorm.DB) *userSeeder {
	return &userSeeder{conn: db}
}

func (ags *userSeeder) Seeder() {
	pass, err := security.Hash("testuser")
	if err != nil {
		return
	}

	ags.create(entity.User{
		Name:     "test1",
		Email:    "user1@example.com",
		Password: pass,
	})
	ags.create(entity.User{
		Name:     "test2",
		Email:    "user2@example.com",
		Password: pass,
	})
	ags.create(entity.User{
		Name:     "test3",
		Email:    "user3@example.com",
		Password: pass,
	})
}

func (ags *userSeeder) create(entity entity.User) error {
	tx := ags.conn.Begin()
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
