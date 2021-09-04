package entity

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type TaskChild struct {
	ID             uint          `json:"id"`
	TaskID         uint          `json:"taskId" gorm:"not null"`
	Task           *Task         `json:"task"`
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

func (tc *TaskChild) TableName() string {
	return "task_children"
}

func (tc *TaskChild) BeforeSave(tx *gorm.DB) error {
	if ok := tc.isExistsTask(tx); !ok {
		return errors.New("親タスクが存在しません。")
	}
	if ok := tc.isExistsTaskStatus(tx); !ok {
		return errors.New("タスクの状態が存在しません。")
	}
	if ok := tc.isExistsTaskPriority(tx); !ok {
		return errors.New("優先順位が存在しません。")
	}
	if ok := tc.isExistsCreateUser(tx); !ok {
		return errors.New("作成者が存在しません。")
	}
	if ok := tc.isExistsAssignUser(tx); !ok {
		return errors.New("割り当てられたユーザーが存在しません。")
	}
	return nil
}

func (tc *TaskChild) isExistsTask(tx *gorm.DB) bool {
	task := &Task{}
	if err := tx.Where("id = ?", tc.TaskID).First(task).Error; err != nil {
		return false
	}
	return true
}

func (tc *TaskChild) isExistsTaskStatus(tx *gorm.DB) bool {
	taskStatus := &TaskStatus{}
	if err := tx.Where("id = ?", tc.TaskStatusID).First(taskStatus).Error; err != nil {
		return false
	}
	return true
}

func (tc *TaskChild) isExistsTaskPriority(tx *gorm.DB) bool {
	taskPriority := &TaskPriority{}
	if err := tx.Where("id = ?", tc.TaskPriorityID).First(taskPriority).Error; err != nil {
		return false
	}
	return true
}

func (tc *TaskChild) isExistsCreateUser(tx *gorm.DB) bool {
	user := &User{}
	if err := tx.Where("id = ?", tc.CreateUserID).First(user).Error; err != nil {
		return false
	}
	return true
}

func (tc *TaskChild) isExistsAssignUser(tx *gorm.DB) bool {
	user := &User{}
	if err := tx.Where("id = ?", tc.AssignUserID).First(user).Error; err != nil {
		return false
	}
	return true
}
