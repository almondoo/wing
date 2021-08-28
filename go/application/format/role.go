package format

import (
	"wing/domain/entity"
)

func RolesFormat(roles []*entity.Role) interface{} {
	format := make([]interface{}, len(roles))
	for i, role := range roles {
		format[i] = map[string]interface{}{
			"id":   role.ID,
			"name": role.Name,
		}
	}
	return format
}

func RoleDetailFormat(role *entity.Role) interface{} {
	return map[string]interface{}{
		"id":   role.ID,
		"name": role.Name,
	}
}
