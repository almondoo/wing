package entity

import (
	"errors"

	"gorm.io/gorm"
)

type TaskStatus struct {
	ID           uint        `json:"id"`
	Name         string      `json:"name" gorm:"size:30;not null;unique"`
	Tasks        []Task      `json:"tasks"`
	TaskChildren []TaskChild `json:"taskChildren"`
}

func (r *TaskStatus) TableName() string {
	return "task_statuses"
}

func (r *TaskStatus) BeforeSave(tx *gorm.DB) (err error) {
	if ok := r.isExistsName(tx); ok {
		return errors.New("既に同じnameが存在しています。")
	}
	return
}

func (r *TaskStatus) isExistsName(tx *gorm.DB) bool {
	role := &Role{}
	if err := tx.Where("name = ?", r.Name).First(role).Error; err != nil {
		return false
	}
	return true
}
