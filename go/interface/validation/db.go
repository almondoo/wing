package validation

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
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

// 下記は使ってないかも
// 調査が必要
type validatorWithDB struct {
	db *gorm.DB
}

type ValidatorWithDB interface {
	IsExists(table string, name string, value interface{}, entity interface{}) bool
}

func NewValidatorWithDB(db *gorm.DB) ValidatorWithDB {
	return &validatorWithDB{db: db}
}

func (d *validatorWithDB) IsExists(table string, name string, value interface{}, entity interface{}) bool {
	if err := d.db.Table(table).Where(name+" = ?", value).Find(&entity).Error; err != nil {
		return false
	}
	return true
}
