package validation

import (
	"github.com/go-playground/validator/v10"
)

//- ユーザーログイン時のリクエスト
type UserLoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

func UserLoginMessage(err error) map[string]string {
	var errorMessages = make(map[string]string)
	errors := err.(validator.ValidationErrors)
	if len(errors) != 0 {
		for i := range errors {
			tag := errors[i].Tag()

			switch errors[i].StructField() {
			case "Email":
				switch tag {
				case "required":
					errorMessages["email"] = "必須項目です。"
				case "email":
					errorMessages["email"] = "メールアドレスの形式で入力してください。"
				}
			case "Password":
				switch tag {
				case "required":
					errorMessages["password"] = "必須項目です。"
				}
			}
		}
		return errorMessages
	}

	return nil
}

//- ユーザー作成時のリクエスト
type UserRegisterRequest struct {
	Name            string `json:"name" form:"name" validate:"required,max=30"`
	Email           string `json:"email" form:"email" validate:"required,email,max=255"`
	Password        string `json:"password" form:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"password_confirm" form:"password_confirm" validate:"required,eqfield=Password"`
}

func UserRegisterMessage(err error) map[string]string {
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

			case "Email":
				switch tag {
				case "required":
					errorMessages["email"] = "必須項目です。"
				case "email":
					errorMessages["email"] = "メールアドレスの形式で入力してください。"
				case "max":
					errorMessages["email"] = "255文字以内で入力してください。"
				}

			case "Password":
				switch tag {
				case "required":
					errorMessages["password"] = "必須項目です。"
				}

			case "PasswordConfirm":
				switch tag {
				case "required":
					errorMessages["password_confirm"] = "必須項目です。"
				case "eqfield":
					errorMessages["password_confirm"] = "パスワードと同じ値を入力してください。"
				}
			}
		}
		return errorMessages
	}

	return nil
}

type UserEditRequest struct {
	Name  string `json:"name" form:"name" validate:"required,max=30"`
	Email string `json:"email" form:"email" validate:"required,email,max=255"`
}

func UserEditMessage(err error) map[string]string {
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

			case "Email":
				switch tag {
				case "required":
					errorMessages["email"] = "必須項目です。"
				case "email":
					errorMessages["email"] = "メールアドレスの形式で入力してください。"
				case "max":
					errorMessages["email"] = "255文字以内で入力してください。"
				}
			}
		}
		return errorMessages
	}

	return nil
}
