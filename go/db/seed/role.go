package seed

import (
	"wing/domain/entity"

	"gorm.io/gorm"
)

type roleSeeder struct {
	conn *gorm.DB
}

func NewRoleSeeder(db *gorm.DB) *roleSeeder {
	return &roleSeeder{conn: db}
}

// 開発者(developer) or オーナー(owner) or 管理者(administrator):全ての権限
// 編集者(editor):作成,更新,参照権限
// 閲覧者(viewer):参照権限
func (rs *roleSeeder) Seeder() {
	rs.create(entity.Role{
		ID:   1,
		Name: "developer",
	})
	rs.create(entity.Role{
		ID:   2,
		Name: "administrator",
	})
	rs.create(entity.Role{
		ID:   3,
		Name: "editor",
	})
	rs.create(entity.Role{
		ID:   4,
		Name: "viewer",
	})
}

func (rs *roleSeeder) create(entity entity.Role) error {
	tx := rs.conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&entity).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
