package format

import (
	"wing/domain/entity"
)

func ProjectsFormat(projects []*entity.Project) interface{} {
	format := make([]interface{}, len(projects))
	for i, project := range projects {
		format[i] = map[string]interface{}{
			"id":      project.ID,
			"name":    project.Name,
			"content": project.Content,
		}
	}
	return format
}

func ProjectDetailFormat(project *entity.Project) interface{} {
	return map[string]interface{}{
		"id":      project.ID,
		"name":    project.Name,
		"content": project.Content,
	}
}
