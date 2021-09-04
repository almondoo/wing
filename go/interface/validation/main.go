package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// CustomValidator
type CustomValidator struct {
	validator *validator.Validate
}

// NewValidator
func NewValidator() echo.Validator {
	validate := validator.New()
	// validate.RegisterValidation("is_japan",isJapan) // カスタムバリデーターを追加
	return &CustomValidator{validator: validate}
}

// Validate validate
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// カスタムバリデーター
// func isJapan(fl validator.FieldLevel) bool {  //引数の型、返り値は固定
//     birthPlace := fl.Field().String()
//     if birthPlace == "Japan" {
//         return true
//     }
//     return false
// }
