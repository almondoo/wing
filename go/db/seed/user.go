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

func (us *userSeeder) Seeder() {
	pass, err := security.Hash("testuser")
	if err != nil {
		return
	}

	us.create(entity.User{
		ID: 1,
		// RoleID:   1,
		Name:     "test1",
		Email:    "user1@example.com",
		Password: pass,
	})
	us.create(entity.User{
		ID: 2,
		// RoleID:   1,
		Name:     "test2",
		Email:    "user2@example.com",
		Password: pass,
	})
	us.create(entity.User{
		ID: 3,
		// RoleID:   1,
		Name:     "test3",
		Email:    "user3@example.com",
		Password: pass,
	})
}

func (us *userSeeder) create(entity entity.User) error {
	tx := us.conn.Begin()
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
