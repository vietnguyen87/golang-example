package converter

import (
	"example-service/dto"
	"example-service/internal/helper"
	"example-service/internal/model"
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
		CreatedAt:   helper.TimeToString(input.CreatedAt),
		UpdatedAt:   helper.TimeToString(input.UpdatedAt),
		DeletedAt:   helper.NullTimeToString(input.DeletedAt),
	}
}
