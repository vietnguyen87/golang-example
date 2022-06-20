package dto

type Task struct {
	ID          uint64  `json:"id"`
	Summary     string  `json:"summary"`
	IsCompleted bool    `json:"is_completed"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	DeletedAt   *string `json:"deleted_at"`
}

type GetTasksResponse struct {
	Data []*Task `json:"data"`
}
