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
	return &CustomValidator{validator: validate}
}

// Validate validate
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
