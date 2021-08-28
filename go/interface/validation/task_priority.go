package validation

import (
	"github.com/go-playground/validator/v10"
)

/*
Requestのsuffixはバリデーションしているだけ
Messageのsuffixはバリデーションにかかった内容を返す
*/

// 共通処理
type TaskPriorityRequest struct {
	Name string `json:"name" form:"name" validate:"required,max=30"`
}

func TaskPriorityMessage(err error) map[string]string {
	var errorMessages = make(map[string]string)
	errors := err.(validator.ValidationErrors)
	if len(errors) != 0 {
		for i := range errors {
			tag := errors[i].Tag()

			switch errors[i].StructField() {
			case "Name":
				switch tag {
				case "required":
					errorMessages["name"] = "必須項目です。"
				case "max":
					errorMessages["name"] = "30文字以内で入力してください。"
				}
			}
		}
		return errorMessages
	}

	return nil
}
