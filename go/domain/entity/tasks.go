package entity

import (
	"errors"
	"time"
	"wing/interface/validation"

	"gorm.io/gorm"
)

type Task struct {
	ID             uint          `json:"id"`
	ProjectID      uint32        `json:"projectId" gorm:"not null" validate:"required"`
	Project        *Project      `json:"project"`
	CreateUserID   uint          `json:"create_user_id"`
	CreateUser     *User         `json:"createUser"`
	Title          string        `json:"title" gorm:"size:255"`
	Content        string        `json:"content" gorm:"type:TEXT"`
	TaskStatusID   uint          `json:"taskStatusId" gorm:"tinyint;not null"`
	TaskStatus     *TaskStatus   `json:"taskStatus"`
	AssignUserID   uint          `json:"assignUserId"`
	AssignUser     *User         `json:"assignUser"`
	TaskPriorityID uint          `json:"taskPriorityId" gorm:"tinyint;not null"`
	TaskPriority   *TaskPriority `json:"taskPriority"`
	StartDate      *time.Time    `json:"startDate"`
	EndDate        *time.Time    `json:"endDate"`
	CreatedAt      time.Time     `json:"createdAt" gorm:"not null"`
	UpdatedAt      time.Time     `json:"updatedAt" gorm:"not null"`
}

func (t *Task) TableName() string {
	return "tasks"
}

func (t *Task) BeforeSave(tx *gorm.DB) error {
	v := validation.DBValidatorInit()
	if err := v.Validate(t); err != nil {
		return err
	}
	if ok := t.isExistsProject(tx); !ok {
		return errors.New("projectが存在しません。")
	}
	return nil
}

func (t *Task) isExistsProject(tx *gorm.DB) bool {
	project := &Project{}
	if err := tx.Where("id = ?", t.ProjectID).First(project).Error; err != nil {
		return false
	}
	return true
}
