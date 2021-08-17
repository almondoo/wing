package persistence

import (
	"wing/domain/entity"
	"wing/domain/repository"

	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) repository.UserRepository {
	return &UserRepository{Conn: conn}
}

func (ur *UserRepository) FindByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	if err := ur.Conn.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) FindByID(id uint) (*entity.User, error) {
	user := &entity.User{}
	if err := ur.Conn.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

//- 作成
func (ur *UserRepository) CreateUser(user *entity.User) (*entity.User, error) {
	tx := ur.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return user, tx.Commit().Error
}

//- 更新
func (ur *UserRepository) Update(user *entity.User) (*entity.User, error) {
	tx := ur.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}
	if err := ur.Conn.Save(user).Error; err != nil {
		tx.Rollback()
		return nil, ErrSave
	}

	return user, tx.Commit().Error
}

//- 削除
func (ur *UserRepository) Delete(user *entity.User) error {
	if err := ur.Conn.Delete(&user).Error; err != nil {
		return ErrDelete
	}

	return nil
}
