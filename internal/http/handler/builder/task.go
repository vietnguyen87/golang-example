package builder

import (
	"example-service/dto"
	"example-service/internal/model"
)

func BuildTasksResponse(tasks []*model.Task, total int64, pagination *dto.Pagination) *dto.ListResp[[]*model.Task] {
	return &dto.ListResp[[]*model.Task]{
		Data: tasks,
		Metadata: &dto.Metadata{
			Pagination: pagination,
			Total:      total,
		},
	}
}
