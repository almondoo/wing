package format

import (
	"wing/domain/entity"
)

func UsersFormat(users []*entity.User) interface{} {
	format := make([]interface{}, len(users))
	for i, user := range users {
		format[i] = map[string]interface{}{
			"id": user.ID,
		}
	}
	return format
}

func UserDetailFormat(user *entity.User) interface{} {
	return map[string]interface{}{
		"id": user.ID,
	}
}
