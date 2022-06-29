package repository

import (
	"context"
	"example-service/internal/constants"
	"example-service/internal/helper"
	"example-service/internal/model"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

//go:generate mockery --name=TaskRepository --case=snake
type TaskRepository interface {
	Find(ctx context.Context, query *model.Query) ([]*model.Task, int64, error)
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepositoryImpl{
		db: db,
	}
}

type taskRepositoryImpl struct {
	db *gorm.DB
}

func (i *taskRepositoryImpl) Find(ctx context.Context, query *model.Query) (tasks []*model.Task, total int64, err error) {
	tx := i.db.WithContext(ctx).Model(&model.Task{})
	fields, values := helper.BuildFilters(query.Q, query.Filters)

	tx = tx.Where(strings.Join(fields, " AND "), values...)
	if query.Sort != nil {
		tx.Order(fmt.Sprintf("%s %s", query.Sort.Key, query.Sort.SortBy))
	}

	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if query.Pagination == nil {
		query.Pagination = helper.BuildPagination(constants.PAGINATION_PAGE, constants.PAGINATION_LIMIT)
	}
	err = tx.Offset(query.Pagination.Offset).Limit(query.Pagination.Limit).Find(&tasks).Error
	return tasks, total, err
}
