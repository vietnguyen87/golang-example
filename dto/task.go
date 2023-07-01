package dto

import "example-service/internal/model"

type Results ListResp[[]*model.Task]

type Task struct {
	ID          uint64  `json:"id"`
	Summary     string  `json:"summary"`
	IsCompleted bool    `json:"is_completed"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	DeletedAt   *string `json:"deleted_at"`
}

type GetResponse struct {
	Tasks      []*Task     `json:"tasks"`
	Total      int64       `json:"total"`
	Pagination *Pagination `json:"pagination"`
}
