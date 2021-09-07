package seed

// import (
// 	"wing/domain/entity"

// 	"gorm.io/gorm"
// )

// type userHaveRoleSeeder struct {
// 	conn *gorm.DB
// }

// func NewUserHaveRoleSeeder(db *gorm.DB) *userHaveRoleSeeder {
// 	return &userHaveRoleSeeder{conn: db}
// }

// func (rs *userHaveRoleSeeder) Seeder() {
// 	rs.create(entity.UserHaveRole{
// 		ID:     1,
// 		UserID: 1,
// 		RoleID: 1,
// 	})
// 	rs.create(entity.UserHaveRole{
// 		ID:     2,
// 		UserID: 2,
// 		RoleID: 6,
// 	})
// 	rs.create(entity.UserHaveRole{
// 		ID:     3,
// 		UserID: 3,
// 		RoleID: 8,
// 	})
// }

// func (rs *userHaveRoleSeeder) create(entity entity.UserHaveRole) error {
// 	tx := rs.conn.Begin()
// 	defer func() {
// 		if r := recover(); r != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	if err := tx.Error; err != nil {
// 		return err
// 	}

// 	if err := tx.Create(&entity).Error; err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	return tx.Commit().Error
// }
