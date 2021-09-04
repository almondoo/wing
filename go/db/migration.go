package db

import (
	"wing/domain/entity"
)

/*
Migration規則
関数名はprefixでmigrateの後に日付を入れる
*/

func (r *Repositories) migrate20210831() error {
	// カラムが確認
	if ok := r.DB.Migrator().HasColumn(&entity.User{}, "role_id"); ok {
		// foreign keyを確認
		if ok := r.DB.Migrator().HasConstraint(&entity.User{}, "fk_roles_users"); ok {
			// foreign key削除
			r.DB.Migrator().DropConstraint(&entity.User{}, "fk_roles_users")
		}
		// カラム削除
		return r.DB.Migrator().DropColumn(&entity.User{}, "role_id")
	}
	return nil
}
