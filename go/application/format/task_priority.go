package format

import (
	"wing/domain/entity"
)

func TaskPrioritiesFormat(taskStatuses []*entity.TaskPriority) interface{} {
	format := make([]interface{}, len(taskStatuses))
	for i, taskStatus := range taskStatuses {
		format[i] = map[string]interface{}{
			"id":   taskStatus.ID,
			"name": taskStatus.Name,
		}
	}
	return format
}

func TaskPriorityDetailFormat(taskStatus *entity.TaskPriority) interface{} {
	return map[string]interface{}{
		"id":   taskStatus.ID,
		"name": taskStatus.Name,
	}
}
