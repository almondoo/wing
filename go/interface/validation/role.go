package validation

import (
	"github.com/go-playground/validator/v10"
)

// 権限作成
type RoleCreateRequest struct {
	// ID   uint   `json:"id" form:"id" validate:"required"`
	Name string `json:"name" form:"name" validate:"required,max=30"`
}

func RoleCreateMessage(err error) map[string]string {
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

// 権限更新
type RoleUpdateRequest struct {
	ID   uint   `json:"id" form:"id" validate:"required"`
	Name string `json:"name" form:"name" validate:"required,max=30"`
}

func RoleUpdateMessage(err error) map[string]string {
	var errorMessages = make(map[string]string)
	errors := err.(validator.ValidationErrors)
	if len(errors) != 0 {
		for i := range errors {
			tag := errors[i].Tag()

			switch errors[i].StructField() {
			case "ID":
				switch tag {
				case "required":
					errorMessages["id"] = "必須項目です。"
				}
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

// 権限削除
type RoleDeleteRequest struct {
	ID uint `json:"id" form:"id" validate:"required"`
}

func RoleDeleteMessage(err error) map[string]string {
	var errorMessages = make(map[string]string)
	errors := err.(validator.ValidationErrors)
	if len(errors) != 0 {
		for i := range errors {
			tag := errors[i].Tag()

			switch errors[i].StructField() {
			case "ID":
				switch tag {
				case "required":
					errorMessages["id"] = "必須項目です。"
				}
			}
		}
		return errorMessages
	}

	return nil
}
