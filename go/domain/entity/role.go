package entity

import (
	"errors"

	"gorm.io/gorm"
)

type Role struct {
	ID           uint           `json:"id"`
	Name         string         `json:"name" gorm:"size:30;not null"`
	Actor        string         `json:"actor" gorm:"size:30;not null"`
	UserHaveRole []UserHaveRole `json:"userHaveRole"`
}

func (r *Role) TableName() string {
	return "roles"
}

func (r *Role) BeforeSave(tx *gorm.DB) (err error) {
	if ok := r.isExistsRecord(tx); ok {
		return errors.New("既に同じ値のレコードが存在しています。")
	}
	return
}

func (r *Role) isExistsRecord(tx *gorm.DB) bool {
	role := &Role{}
	if err := tx.Where("name = ? AND actor = ?", r.Name, r.Actor).First(role).Error; err != nil {
		return false
	}
	return true
}
