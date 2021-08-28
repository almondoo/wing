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
	TaskChildren   []TaskChild   `json:"taskChildren"`
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
		return errors.New("プロジェクトが存在しません。")
	}
	if ok := t.isExistsTaskStatus(tx); !ok {
		return errors.New("タスクの状態が存在しません。")
	}
	if ok := t.isExistsTaskPriority(tx); !ok {
		return errors.New("優先順位が存在しません。")
	}
	if ok := t.isExistsCreateUser(tx); !ok {
		return errors.New("作成者が存在しません。")
	}
	if ok := t.isExistsAssignUser(tx); !ok {
		return errors.New("割り当てられたユーザーが存在しません。")
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

func (t *Task) isExistsTaskStatus(tx *gorm.DB) bool {
	taskStatus := &TaskStatus{}
	if err := tx.Where("id = ?", t.TaskStatusID).First(taskStatus).Error; err != nil {
		return false
	}
	return true
}

func (t *Task) isExistsTaskPriority(tx *gorm.DB) bool {
	taskPriority := &TaskPriority{}
	if err := tx.Where("id = ?", t.TaskPriorityID).First(taskPriority).Error; err != nil {
		return false
	}
	return true
}

func (t *Task) isExistsCreateUser(tx *gorm.DB) bool {
	user := &User{}
	if err := tx.Where("id = ?", t.CreateUserID).First(user).Error; err != nil {
		return false
	}
	return true
}

func (t *Task) isExistsAssignUser(tx *gorm.DB) bool {
	user := &User{}
	if err := tx.Where("id = ?", t.AssignUserID).First(user).Error; err != nil {
		return false
	}
	return true
}
