package dto

type Task struct {
	ID          uint64  `json:"id"`
	Summary     string  `json:"summary"`
	IsCompleted bool    `json:"is_completed"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	DeletedAt   *string `json:"deleted_at"`
}

type GetResponse struct {
	Data       []*Task     `json:"data"`
	Total      int64       `json:"total"`
	Pagination *Pagination `json:"pagination"`
}
