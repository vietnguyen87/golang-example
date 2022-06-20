package converter

import (
	"mapi-service/dto"
	"mapi-service/internal/model"
)

func TasksToDTO(input []*model.Task) []*dto.Task {
	output := make([]*dto.Task, 0)

	for _, v := range input {
		output = append(output, TaskToDTO(v))
	}

	return output
}

func TaskToDTO(input *model.Task) *dto.Task {
	if input == nil {
		return nil
	}

	return &dto.Task{
		ID:          input.ID,
		Summary:     input.Summary,
		IsCompleted: input.IsCompleted,
		CreatedAt:   timeToString(input.CreatedAt),
		UpdatedAt:   timeToString(input.UpdatedAt),
		DeletedAt:   nullTimeToString(input.DeletedAt),
	}
}
