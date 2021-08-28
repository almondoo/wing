package format

import (
	"wing/domain/entity"
)

func TasksFormat(taskStatuses []*entity.Task) interface{} {
	format := make([]interface{}, len(taskStatuses))
	for i, taskStatus := range taskStatuses {
		format[i] = map[string]interface{}{
			"id": taskStatus.ID,
		}
	}
	return format
}

func TaskDetailFormat(taskStatus *entity.Task) interface{} {
	return map[string]interface{}{
		"id": taskStatus.ID,
	}
}
