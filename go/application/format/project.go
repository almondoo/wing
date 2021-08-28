package format

import (
	"wing/domain/entity"
)

func ProjectsFormat(taskStatuses []*entity.Project) interface{} {
	format := make([]interface{}, len(taskStatuses))
	for i, taskStatus := range taskStatuses {
		format[i] = map[string]interface{}{
			"id": taskStatus.ID,
		}
	}
	return format
}

func ProjectDetailFormat(taskStatus *entity.Project) interface{} {
	return map[string]interface{}{
		"id": taskStatus.ID,
	}
}
