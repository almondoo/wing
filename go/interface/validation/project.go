package validation

import (
	"github.com/go-playground/validator/v10"
)

/*
Requestのsuffixはバリデーションしているだけ
Messageのsuffixはバリデーションにかかった内容を返す
*/

// 共通処理
type ProjectRequest struct {
	Name    string `json:"name" form:"name" validate:"required,max=255"`
	Content string `json:"content" from:"content" validate:"required,max=20000"`
}

func ProjectMessage(err error) map[string]string {
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
			case "Content":
				switch tag {
				case "required":
					errorMessages["content"] = "必須項目です。"
				case "max":
					errorMessages["content"] = "20000文字以内で入力してください。"
				}
			}
		}
		return errorMessages
	}

	return nil
}
