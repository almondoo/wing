package format

import (
	"wing/domain/entity"
)

func TaskChildrenFormat(taskStatuses []*entity.TaskChild) interface{} {
	format := make([]interface{}, len(taskStatuses))
	for i, taskStatus := range taskStatuses {
		format[i] = map[string]interface{}{
			"id": taskStatus.ID,
		}
	}
	return format
}

func TaskChildDetailFormat(taskStatus *entity.TaskChild) interface{} {
	return map[string]interface{}{
		"id": taskStatus.ID,
	}
}
