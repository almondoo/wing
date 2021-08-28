package validation

import (
	"github.com/go-playground/validator/v10"
)

type dbValidator struct {
	vali *validator.Validate
}

func DBValidatorInit() *dbValidator {
	return &dbValidator{vali: validator.New()}
}

func (v *dbValidator) Validate(entity interface{}) error {
	return v.vali.Struct(entity)
}
