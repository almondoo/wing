package entity

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              uint `json:"id"`
	Role            *Role
	RoleID          uint32     `json:"roleId" gorm:"not null"`
	Name            string     `json:"name" gorm:"size:100;not null"`
	Email           string     `json:"email" gorm:"size:255;not null;unique"`
	EmailVerifiedAt *time.Time `json:"emailVerifiedAt"`
	Password        string     `json:"-" gorm:"size:255;not null"`
	CreatedAt       time.Time  `json:"createdAt" gorm:"not null"`
	UpdatedAt       time.Time  `json:"updatedAt" gorm:"not null"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if ok := u.isExistsEmail(tx); ok {
		return errors.New("emailが既に存在しています。")
	}

	return
}

func (u *User) isExistsEmail(tx *gorm.DB) bool {
	user := &User{}
	if err := tx.Where("email = ?", u.Email).First(user).Error; err != nil {
		return false
	}
	return true
}
