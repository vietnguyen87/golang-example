package repository

import (
	"context"
	"example-service/internal/model"
	"example-service/pkg/logger"
	"gorm.io/gorm"
)

//go:generate mockery --name=TaskRepository --case=snake
type TaskRepository interface {
	GetTasks(ctx context.Context) ([]*model.Task, error)
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepositoryImpl{
		db: db,
	}
}

type taskRepositoryImpl struct {
	db *gorm.DB
}

func (i *taskRepositoryImpl) GetTasks(ctx context.Context) ([]*model.Task, error) {
	log := logger.CToL(ctx, "GetTasks")

	var tasks []*model.Task
	err := i.db.WithContext(ctx).Find(&tasks).Error
	if err != nil {
		log.WithError(err).Errorf("gorm.DB returns error when .Find: %s", err.Error())
		return nil, err
	}

	return tasks, nil
}
