package persistence

import (
	"wing/domain/entity"
	"wing/domain/repository"

	"gorm.io/gorm"
)

type userHaveRoleRepository struct {
	Conn *gorm.DB
}

func NewUserHaveRoleRepository(conn *gorm.DB) repository.UserHaveRoleRepository {
	return &userHaveRoleRepository{Conn: conn}
}

// FindsByUserID user_idのレコードを取得
func (rr *userHaveRoleRepository) FindsByUserID(userId uint) ([]*entity.UserHaveRole, error) {
	userHaveRoles := []*entity.UserHaveRole{}

	if err := rr.Conn.Where("user_id = ?", userId).Preload("Role").Find(&userHaveRoles).Error; err != nil {
		return nil, err
	}

	return userHaveRoles, nil
}

// FindByConditions 自由に条件を作る
func (rr *userHaveRoleRepository) FindByConditions(conditions map[string]interface{}) (*entity.UserHaveRole, error) {
	var userHaveRole *entity.UserHaveRole
	if err := rr.Conn.Where(conditions).Find(&userHaveRole).Error; err != nil {
		return nil, err
	}
	return userHaveRole, nil
}

// Create 作成
func (rr *userHaveRoleRepository) Create(userHaveRole *entity.UserHaveRole) (*entity.UserHaveRole, error) {
	tx := rr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Create(&userHaveRole).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return userHaveRole, tx.Commit().Error
}

// Update 更新
func (rr *userHaveRoleRepository) Update(userHaveRole *entity.UserHaveRole) (*entity.UserHaveRole, error) {
	tx := rr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Model(&userHaveRole).Updates(&userHaveRole).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return userHaveRole, tx.Commit().Error
}

// Delete 削除
func (rr *userHaveRoleRepository) Delete(userId uint, roleId uint) error {
	tx := rr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Where("user_id = ? AND role_id = ?", userId, roleId).Delete(&entity.UserHaveRole{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
