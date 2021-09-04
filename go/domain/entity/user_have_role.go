package entity

import (
	"errors"

	"gorm.io/gorm"
)

type UserHaveRole struct {
	ID     uint `json:"id"`
	UserID uint `json:"user_id"`
	User   User `json:"user"`
	RoleID uint `json:"role_id"`
	Role   Role `json:"role"`
}

func (r *UserHaveRole) TableName() string {
	return "user_have_roles"
}

// BeforeSave 作成もしくは更新する前に実行される
func (uhr *UserHaveRole) BeforeSave(tx *gorm.DB) error {
	if ok := uhr.isExistsRole(tx); !ok {
		return errors.New("権限が存在しません。")
	}
	return nil
}

// BeforeCreate 作成する前に実行される
func (uhr *UserHaveRole) BeforeCreate(tx *gorm.DB) error {
	if ok := uhr.isExistsRoleName(tx); ok {
		return errors.New("既に同じ名前の権限は付与されています。")
	}
	return nil
}

// isExistsRole 権限があるか
func (uhr *UserHaveRole) isExistsRole(tx *gorm.DB) bool {
	role := &Role{}
	if err := tx.First(role, uhr.RoleID).Error; err != nil {
		return false
	}
	return true
}

// isExistsRoleName 権限の同じ名前があるか
func (uhr *UserHaveRole) isExistsRoleName(tx *gorm.DB) bool {
	role := &Role{}
	tx.First(role, uhr.RoleID)

	userHaveRole := &UserHaveRole{}
	if err := tx.Where("user_id = ?", uhr.UserID).Preload("Role").First(userHaveRole).Error; err != nil {
		return false
	}
	if role.Name == userHaveRole.Role.Name {
		return true
	}
	return false
}
