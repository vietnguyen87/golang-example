package model

type Task struct {
	Model
	Summary     string
	IsCompleted bool `json:"isCompleted" select:"is_completed"`
}
