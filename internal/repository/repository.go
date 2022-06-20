package repository

import "gorm.io/gorm"

//go:generate mockery --name=Repository --case=snake
type Repository interface {
	TaskRepository() TaskRepository
}

func New(db *gorm.DB) Repository {
	return &repositoryImpl{
		taskRepository: NewTaskRepository(db),
	}
}

type repositoryImpl struct {
	taskRepository TaskRepository
}

func (i *repositoryImpl) TaskRepository() TaskRepository {
	return i.taskRepository
}
