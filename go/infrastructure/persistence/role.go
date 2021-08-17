package persistence

import (
	"wing/domain/entity"
	"wing/domain/repository"

	"gorm.io/gorm"
)

type roleRepository struct {
	Conn *gorm.DB
}

func NewRoleRepository(conn *gorm.DB) repository.RoleRepository {
	return &roleRepository{Conn: conn}
}

func (rr *roleRepository) CreateRole(product *entity.Role) error {
	tx := rr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&product).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (rr *roleRepository) FindByEmail(email string) (entity.Role, error) {
	var product entity.Role
	rr.Conn.Where("email = ?", email).First(&product)
	return product, nil
}

//- IDで取得
func (rr *roleRepository) FindByID(id uint) (*entity.Role, error) {
	product := &entity.Role{ID: id}

	if err := rr.Conn.First(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

//- 作成
func (rr *roleRepository) Create(product *entity.Role) (*entity.Role, error) {
	if err := rr.Conn.Create(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

//- 更新
func (rr *roleRepository) Update(product *entity.Role) (*entity.Role, error) {
	if err := rr.Conn.Model(&product).Updates(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

//- 削除
func (rr *roleRepository) Delete(product *entity.Role) error {
	if err := rr.Conn.Delete(&product).Error; err != nil {
		return err
	}

	return nil
}
