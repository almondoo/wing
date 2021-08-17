package entity

import (
	"errors"
	"wing/interface/validation"

	"gorm.io/gorm"
)

type Role struct {
	ID    uint   `json:"id"`
	Name  string `json:"name" gorm:"size:30;not null;unique" validate:"required,max=30"`
	Users []User `json:"users"`
}

func (r *Role) TableName() string {
	return "roles"
}

func (r *Role) BeforeSave(tx *gorm.DB) (err error) {
	v := validation.DBValidatorInit()
	err = v.Validate(r)
	if err != nil {
		return err
	}
	if ok := r.isExistsName(tx); ok {
		return errors.New("既に同じnameが存在しています。")
	}
	return
}

func (r *Role) isExistsName(tx *gorm.DB) bool {
	role := &Role{}
	if err := tx.Where("name = ?", r.Name).First(role).Error; err != nil {
		return false
	}
	return true
}
