package entity

import (
	"errors"
	"wing/interface/validation"

	"gorm.io/gorm"
)

type TaskStatus struct {
	ID           uint        `json:"id"`
	Name         string      `json:"name" gorm:"size:30;not null;unique" validate:"required,max=30"`
	Tasks        []Task      `json:"tasks"`
	TaskChildren []TaskChild `json:"taskChildren"`
}

func (r *TaskStatus) TableName() string {
	return "task_statuses"
}

func (r *TaskStatus) BeforeSave(tx *gorm.DB) (err error) {
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

func (r *TaskStatus) isExistsName(tx *gorm.DB) bool {
	role := &Role{}
	if err := tx.Where("name = ?", r.Name).First(role).Error; err != nil {
		return false
	}
	return true
}
