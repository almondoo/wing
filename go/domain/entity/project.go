package entity

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID        uint32    `json:"id"`
	Name      string    `json:"name" gorm:"size:255;not null;unique"`
	Content   string    `json:"content" gorm:"size:text"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null"`
	Tasks     []Task    `json:"tasks"`
}

func (p *Project) TableName() string {
	return "projects"
}

func (p *Project) BeforeSave(tx *gorm.DB) (err error) {
	if ok := p.isExistsName(tx); ok {
		return errors.New("既に同じプロジェクト名が存在しています。")
	}
	return
}

func (p *Project) isExistsName(tx *gorm.DB) bool {
	project := &Project{}
	if err := tx.Where("name = ?", p.Name).First(project).Error; err != nil {
		return false
	}
	return true
}
