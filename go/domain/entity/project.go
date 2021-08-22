package entity

import (
	"time"
	"wing/interface/validation"

	"gorm.io/gorm"
)

type Project struct {
	ID        uint32    `json:"id"`
	Name      string    `json:"name" gorm:"size:255;not null" validate:"required,max=255"`
	Content   string    `json:"content" gorm:"size:text"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null"`
	Tasks     []Task    `json:"tasks"`
}

func (p *Project) TableName() string {
	return "projects"
}

func (p *Project) BeforeSave(tx *gorm.DB) (err error) {
	v := validation.DBValidatorInit()
	err = v.Validate(p)
	if err != nil {
		return err
	}
	return
}
