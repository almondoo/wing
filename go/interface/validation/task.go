package validation

import (
	"time"

	"github.com/go-playground/validator/v10"
)

/*
Requestのsuffixはバリデーションしているだけ
Messageのsuffixはバリデーションにかかった内容を返す
*/

// TaskRequest 共通処理
type TaskRequest struct {
	ProjectID      uint32    `json:"project_id" form:"project_id" validate:"required,numeric"`
	Title          string    `json:"title" form:"title" validate:"required,max=255"`
	Content        string    `json:"content" form:"content" validate:"max=20000"`
	TaskStatusID   uint      `json:"task_status_id" form:"task_status_id" validate:"required,numeric"`
	AssignUserID   uint      `json:"assign_user_id" form:"assign_user_id" validate:"numeric"`
	TaskPriorityID uint      `json:"task_priority_id" form:"task_priority_id" validate:"numeric"`
	StartDate      time.Time `json:"start_date" form:"start_date" validate:"datetime"`
	EndDate        time.Time `json:"end_date" form:"end_date" validate:"datetime"`
}

// TaskMessage メッセージ
func TaskMessage(err error) map[string]string {
	var errorMessages = make(map[string]string)
	errors := err.(validator.ValidationErrors)
	if len(errors) != 0 {
		for i := range errors {
			tag := errors[i].Tag()

			switch errors[i].StructField() {
			case "ProjectID":
				switch tag {
				case "required":
					errorMessages["project_id"] = "必須項目です。"
				case "max":
					errorMessages["project_id"] = "30文字以内で入力してください。"
				}
			case "Title":
				switch tag {
				case "required":
					errorMessages["title"] = "必須項目です。"
				case "max":
					errorMessages["title"] = "255文字以内で入力してください。"
				}
			case "Content":
				switch tag {
				case "max":
					errorMessages["content"] = "20000文字以内で入力してください。"
				}
			case "TaskStatusID":
				switch tag {
				case "numeric":
					errorMessages["task_status_id"] = "必須項目です。"
				}
			case "AssignUserID":
				switch tag {
				case "numeric":
					errorMessages["assign_user_id"] = "必須項目です。"
				}
			case "TaskPriorityID":
				switch tag {
				case "numeric":
					errorMessages["task_priority_id"] = "必須項目です。"
				}
			case "StartDate":
				switch tag {
				case "datetime":
					errorMessages["start_date"] = "正しいフォーマットにしてください。"
				}
			case "EndDate":
				switch tag {
				case "datetime":
					errorMessages["end_date"] = "正しいフォーマットにしてください。"
				}
			}
		}
		return errorMessages
	}

	return nil
}
