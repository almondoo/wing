package persistence

import (
	"wing/domain/entity"
	"wing/domain/repository"

	"gorm.io/gorm"
)

type roleRepository struct {
	Conn *gorm.DB
}

func NewRoleRepository(conn *gorm.DB) repository.RoleRepository {
	return &roleRepository{Conn: conn}
}

// IDで取得
func (rr *roleRepository) FindByID(id uint) (*entity.Role, error) {
	role := &entity.Role{ID: id}

	if err := rr.Conn.First(&role).Error; err != nil {
		return nil, err
	}

	return role, nil
}

// Create 作成
func (rr *roleRepository) Create(role *entity.Role) (*entity.Role, error) {
	tx := rr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Create(&role).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return role, tx.Commit().Error
}

// Update 更新
func (rr *roleRepository) Update(role *entity.Role) (*entity.Role, error) {
	tx := rr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Model(&role).Updates(&role).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return role, tx.Commit().Error
}

// Delete 削除
func (rr *roleRepository) Delete(id uint) error {
	tx := rr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Delete(&entity.Role{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
