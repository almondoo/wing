package format

import (
	"wing/domain/entity"
)

func TaskStatusesFormat(taskStatuses []*entity.TaskStatus) interface{} {
	format := make([]interface{}, len(taskStatuses))
	for i, taskStatus := range taskStatuses {
		format[i] = map[string]interface{}{
			"id":   taskStatus.ID,
			"name": taskStatus.Name,
		}
	}
	return format
}

func TaskStatusDetailFormat(taskStatus *entity.TaskStatus) interface{} {
	return map[string]interface{}{
		"id":   taskStatus.ID,
		"name": taskStatus.Name,
	}
}
